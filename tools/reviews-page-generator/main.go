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
	"github.com/yuin/goldmark/renderer/html"
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
	commentsRe = regexp.MustCompile(`\{\{#comments\}\}[\s\S]*?\{\{/comments\}\}`)
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
	Avatar    string `json:"avatar"`
	Timestamp string `json:"timestamp"`
	Body      string `json:"body"`
}

func initPaths() {
	baseDir, _ = os.Getwd()
	jsonDir = filepath.Join(baseDir, "comments_json")
	imgsDir = filepath.Join(baseDir, "downloaded_images")
	tmplPath = filepath.Join(baseDir, "template.mustache")
	outPath = filepath.Join(baseDir, "reviews.html")
	dataPath = filepath.Join(baseDir, "reviews-data.json")
	pubDir = filepath.Join(baseDir, "..", "..", "public")
}

func setupDirs() {
	for _, d := range []string{jsonDir, imgsDir} {
		os.MkdirAll(d, 0755)
	}
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
		os.WriteFile(path, data, 0644)
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

func downloadImages(client *http.Client, comments []ghComment) {
	var (
		mu   sync.Mutex
		wg   sync.WaitGroup
		seen = make(map[string]bool)
	)
	sem := make(chan struct{}, concurrency)

	for _, c := range comments {
		urls := imageURLRe.FindAllString(c.Body, -1)
		for _, u := range urls {
			u = strings.Split(u, "?")[0]
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

				name := filepath.Base(url)
				path := filepath.Join(imgsDir, name)
				if _, err := os.Stat(path); err == nil {
					return
				}

				resp, err := client.Get(url)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					return
				}
				os.WriteFile(path, data, 0644)
				fmt.Printf("  Downloaded image: %s\n", name)
			}(u)
		}
	}
	wg.Wait()
}

func prepareComments(comments []ghComment) []reviewEntry {
	md := goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))
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
		for _, url := range imageURLRe.FindAllString(bodyHTML, -1) {
			name := filepath.Base(strings.Split(url, "?")[0])
			bodyHTML = strings.ReplaceAll(bodyHTML, url, "./downloaded_images/"+name)
		}

		entries = append(entries, reviewEntry{
			URL:       c.HTMLURL,
			Author:    c.User.Login,
			Avatar:    "",
			Timestamp: timestamp,
			Body:      bodyHTML,
		})
	}

	return entries
}

func render(entries []reviewEntry) {
	tmpl, err := os.ReadFile(tmplPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading template: %v\n", err)
		os.Exit(1)
	}

	output := string(tmpl)
	output = commentsRe.ReplaceAllString(output, "")
	output = strings.ReplaceAll(output, "{{ total_comments }}", fmt.Sprintf("%d", len(entries)))

	os.WriteFile(outPath, []byte(output), 0644)
	fmt.Printf("  Generated: %s\n", outPath)

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.Encode(entries)
	data := bytes.TrimRight(b.Bytes(), "\n")
	os.WriteFile(dataPath, data, 0644)
	fmt.Printf("  Generated: %s\n", dataPath)
}

func copyToPublic() {
	copyDir(imgsDir, filepath.Join(pubDir, "downloaded_images"))
	copyFile(outPath, filepath.Join(pubDir, "reviews.html"))
	copyFile(dataPath, filepath.Join(pubDir, "reviews-data.json"))
	fmt.Println("  Copied assets → public/")
}

func copyDir(src, dst string) {
	entries, err := os.ReadDir(src)
	if err != nil {
		return
	}
	os.MkdirAll(dst, 0755)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		srcPath := filepath.Join(src, e.Name())
		dstPath := filepath.Join(dst, e.Name())
		data, err := os.ReadFile(srcPath)
		if err != nil {
			continue
		}
		os.WriteFile(dstPath, data, 0644)
	}
}

func copyFile(src, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  Error copying %s: %v\n", src, err)
		return
	}
	os.WriteFile(dst, data, 0644)
}

func main() {
	forceFetch := flag.Bool("force-fetch", false, "Ignore cached JSON, re-fetch from GitHub")
	flag.Parse()

	initPaths()
	setupDirs()

	client := &http.Client{Timeout: 60 * time.Second}

	fmt.Println("Fetching comments...")
	comments := fetchComments(client, *forceFetch)
	fmt.Printf("Processing %d comments...\n", len(comments))

	downloadImages(client, comments)
	entries := prepareComments(comments)

	fmt.Println("Rendering template...")
	render(entries)
	copyToPublic()
	fmt.Println("Done.")
}
