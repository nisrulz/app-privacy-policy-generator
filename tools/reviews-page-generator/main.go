package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
)

const (
	owner       = "nisrulz"
	repoName    = "app-privacy-policy-generator"
	issueNumber = "65"
	apiURL      = "https://api.github.com/repos/" + owner + "/" + repoName + "/issues/" + issueNumber + "/comments"
	concurrency = 8
	oneWeek     = 7 * 24 * time.Hour
)

var (
	baseDir, jsonDir, imgsDir           string
	tmplPath, outPath, dataPath, pubDir string

	imageURLRe = regexp.MustCompile(`https?://[^\s]+\.(?:png|jpg|jpeg|gif)`)
	ghAssetRe  = regexp.MustCompile(`https://github\.com/user-attachments/assets/[a-f0-9-]+`)
	commentsRe = regexp.MustCompile(`\{\{#comments\}\}[\s\S]*?\{\{/comments\}\}`)

	imageMap map[string]string
)

type ghUser struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type ghComment struct {
	HTMLURL   string `json:"html_url"`
	User      ghUser `json:"user"`
	CreatedAt string `json:"created_at"`
	Body      string `json:"body"`
}

type reviewEntry struct {
	URL       string `json:"url"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
	Body      string `json:"body"`
}

func initPaths() error {
	baseDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}
	jsonDir = filepath.Join(baseDir, "comments_json")
	imgsDir = filepath.Join(baseDir, "downloaded_images")
	tmplPath = filepath.Join(baseDir, "template.mustache")
	pubDir = filepath.Join(baseDir, "..", "..", "public")
	outPath = filepath.Join(pubDir, "reviews.html")
	dataPath = filepath.Join(pubDir, "reviews-data.json")
	return nil
}

func setupDirs() error {
	for _, d := range []string{jsonDir, imgsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("create %s: %w", d, err)
		}
	}
	return nil
}

func fetchComments(client *http.Client, force bool) []ghComment {
	var all []ghComment
	page := 1

	for {
		path := filepath.Join(jsonDir, fmt.Sprintf("comments_page_%d.json", page))

		if !force {
			if fi, err := os.Stat(path); err == nil && time.Since(fi.ModTime()) < oneWeek {
				data, err := os.ReadFile(path)
				if err == nil {
					var pageComments []ghComment
					if json.Unmarshal(data, &pageComments) == nil {
						fmt.Printf("  Loaded cached: comments_page_%d.json\n", page)
						all = append(all, pageComments...)
						if len(pageComments) == 0 {
							break
						}
						page++
						continue
					}
				}
			}
		}

		pageComments := fetchPage(client, page)
		if pageComments == nil {
			break
		}
		data, _ := json.Marshal(pageComments)
		if err := os.WriteFile(path, data, 0644); err != nil {
			fmt.Printf("  Warning: could not cache %s: %v\n", path, err)
		}
		fmt.Printf("  Fetched comments_page_%d.json (%d comments)\n", page, len(pageComments))
		time.Sleep(time.Second)

		all = append(all, pageComments...)
		if len(pageComments) == 0 {
			break
		}
		page++
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt > all[j].CreatedAt
	})
	return all
}

func fetchPage(client *http.Client, page int) []ghComment {
	for attempt := 0; attempt < 3; attempt++ {
		req, _ := http.NewRequest("GET", apiURL, nil)
		q := req.URL.Query()
		q.Set("per_page", "100")
		q.Set("page", fmt.Sprintf("%d", page))
		req.URL.RawQuery = q.Encode()
		req.Header.Set("User-Agent", repoName)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("  Error fetching page %d: %v\n", page, err)
			return nil
		}

		if resp.StatusCode == http.StatusForbidden {
			resp.Body.Close()
			fmt.Println("  Rate limited, waiting 60s...")
			time.Sleep(60 * time.Second)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("  Error %d on page %d\n", resp.StatusCode, page)
			return nil
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var comments []ghComment
		if err := json.Unmarshal(body, &comments); err != nil {
			fmt.Printf("  Error parsing page %d: %v\n", page, err)
			return nil
		}
		return comments
	}
	return nil
}

func contentExt(ct string) string {
	ct = strings.TrimSpace(strings.Split(ct, ";")[0])
	switch ct {
	case "image/png":
		return ".png"
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	default:
		return ".png"
	}
}

// findExistingAsset resolves an asset UUID to an already-downloaded filename.
// Asset URLs carry no extension, so the on-disk form is discovered by
// globbing. Selection must not depend on glob ordering, which is lexical and
// therefore arbitrarily picks .png over .webp when both exist; prefer the
// format public/ actually serves and fall back to lexical order only to
// break ties among the remaining formats.
func findExistingAsset(name string) (string, bool) {
	matches, err := filepath.Glob(filepath.Join(imgsDir, name+".*"))
	if err != nil || len(matches) == 0 {
		return "", false
	}
	sort.Strings(matches)
	for _, m := range matches {
		if strings.EqualFold(filepath.Ext(m), ".webp") {
			return filepath.Base(m), true
		}
	}
	return filepath.Base(matches[0]), true
}

func downloadImages(client *http.Client, comments []ghComment) {
	imageMap = make(map[string]string)

	var (
		mu   sync.Mutex
		wg   sync.WaitGroup
		seen = make(map[string]bool)
	)
	sem := make(chan struct{}, concurrency)

	for _, c := range comments {
		var urls []string
		for _, u := range imageURLRe.FindAllString(c.Body, -1) {
			urls = append(urls, strings.Split(u, "?")[0])
		}
		for _, u := range ghAssetRe.FindAllString(c.Body, -1) {
			urls = append(urls, u)
		}
		for _, u := range urls {
			mu.Lock()
			if seen[u] {
				mu.Unlock()
				continue
			}
			seen[u] = true
			mu.Unlock()

			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()

				isAsset := ghAssetRe.MatchString(url)
				name := filepath.Base(url)
				path := filepath.Join(imgsDir, name)
				if isAsset {
					if existing, ok := findExistingAsset(name); ok {
						mu.Lock()
						imageMap[url] = existing
						mu.Unlock()
						return
					}
				} else if _, err := os.Stat(path); err == nil {
					mu.Lock()
					imageMap[url] = name
					mu.Unlock()
					return
				}

				resp, err := client.Get(url)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					return
				}

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					return
				}

				if isAsset {
					ext := contentExt(resp.Header.Get("Content-Type"))
					name = name + ext
					path = filepath.Join(imgsDir, name)
				}

				if err := os.WriteFile(path, data, 0644); err != nil {
					fmt.Printf("  Warning: could not write image %s: %v\n", name, err)
					return
				}
				fmt.Printf("  Downloaded image: %s\n", name)

				mu.Lock()
				imageMap[url] = name
				mu.Unlock()
			}(u)
		}
	}
	wg.Wait()
}

func prepareComments(comments []ghComment) []reviewEntry {
	md := goldmark.New()
	var buf bytes.Buffer

	entries := make([]reviewEntry, 0, len(comments))

	for _, c := range comments {
		t, _ := time.Parse("2006-01-02T15:04:05Z", c.CreatedAt)
		timestamp := t.Format("02 Jan 2006")

		buf.Reset()
		if err := md.Convert([]byte(c.Body), &buf); err != nil {
			buf.WriteString(c.Body)
		}
		bodyHTML := buf.String()

		bodyHTML = strings.ReplaceAll(bodyHTML, c.User.AvatarURL, "")

		for url, name := range imageMap {
			bodyHTML = strings.ReplaceAll(bodyHTML, url, "./downloaded_images/"+name)
		}

		entries = append(entries, reviewEntry{
			URL:       c.HTMLURL,
			Author:    c.User.Login,
			Timestamp: timestamp,
			Body:      bodyHTML,
		})
	}

	return entries
}

func render(entries []reviewEntry) error {
	tmpl, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("read template %s: %w", tmplPath, err)
	}

	output := string(tmpl)
	output = commentsRe.ReplaceAllString(output, "")
	output = strings.ReplaceAll(output, "{{ total_comments }}", fmt.Sprintf("%d", len(entries)))

	if err := os.WriteFile(outPath, []byte(output), 0644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	fmt.Printf("  Generated: %s\n", outPath)

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(entries); err != nil {
		return fmt.Errorf("encode reviews data: %w", err)
	}
	data := bytes.TrimRight(b.Bytes(), "\n")
	if err := os.WriteFile(dataPath, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", dataPath, err)
	}
	fmt.Printf("  Generated: %s\n", dataPath)

	return nil
}

func main() {
	forceFetch := flag.Bool("force-fetch", false, "Ignore cached JSON, re-fetch from GitHub")
	flag.Parse()

	if err := initPaths(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}
	if err := setupDirs(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}

	client := &http.Client{Timeout: 60 * time.Second}

	fmt.Println("Fetching comments...")
	comments := fetchComments(client, *forceFetch)
	fmt.Printf("Processing %d comments...\n", len(comments))

	downloadImages(client, comments)
	entries := prepareComments(comments)

	fmt.Println("Rendering template...")
	if err := render(entries); err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done.")
}
