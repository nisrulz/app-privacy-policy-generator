package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	langFlag := flag.String("lang", "", "Comma-separated list of locales to build (default: all)")
	cleanFlag := flag.Bool("clean", false, "Clean public/ directory before building")
	watchFlag := flag.Bool("watch", false, "Watch for file changes and rebuild")
	serveFlag := flag.Bool("serve", false, "Build and serve locally")
	portFlag := flag.String("port", "8000", "Port to serve on")
	flag.Parse()

	langs, err := determineLangs(*langFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, " ❌  %v\n", err)
		os.Exit(1)
	}

	switch {
	case *serveFlag:
		runServe(langs, *cleanFlag, *portFlag)
	case *watchFlag:
		runWatch(langs)
	default:
		runBuild(langs, *cleanFlag)
	}
}

func runBuild(langs []string, clean bool) {
	start := time.Now()

	fmt.Printf("\n 👨🏻‍💻 Starting build for locales: %s\n\n", strings.Join(langs, ", "))

	if clean {
		step("Cleaning public/", func() error {
			return cleanPublic()
		})
	}

	step("STEP 1: Copy CSS", func() error {
		return buildCSS()
	})

	step("STEP 2: Convert YAML → JS (third-party services)", func() error {
		return buildThirdPartyJS()
	})

	step("STEP 3: Build locales registry JS", func() error {
		return buildLocalesRegistry()
	})

	step("STEP 4: Render Pug → HTML (reference for CSS minification)", func() error {
		return renderPug("en", "public")
	})

	step("STEP 5: Minify JS", func() error {
		return buildMinifyJS()
	})

	step("STEP 6: Copy vendor assets", func() error {
		return copyVendorAssets()
	})

	step("STEP 7: Per-locale builds", func() error {
		return buildLocales(langs)
	})

	step("STEP 8: Cache-busting", func() error {
		return cacheBust()
	})

	elapsed := time.Since(start)
	fmt.Printf("\n 📈  Build complete in %s\n\n", elapsed)
}

func step(name string, fn func() error) {
	start := time.Now()
	if err := fn(); err != nil {
		fmt.Fprintf(os.Stderr, " ❌  %s: %v\n", name, err)
		os.Exit(1)
	}
	fmt.Printf(" ✅  %s (%s)\n", name, time.Since(start).Round(time.Millisecond))
}

func determineLangs(flagLang string) ([]string, error) {
	if flagLang != "" {
		langs := strings.Split(flagLang, ",")
		for i, l := range langs {
			langs[i] = strings.TrimSpace(l)
		}
		return langs, nil
	}
	return getLocales()
}
