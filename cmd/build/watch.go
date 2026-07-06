package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func runWatch(langs []string) {
	fmt.Println(" 👀  Watching for changes...")

	watchDirs := []string{"src", "public"}
	done := make(chan bool)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for _, dir := range watchDirs {
		go watchDir(dir, langs, done)
	}

	<-sigChan
	fmt.Println("\n 🛑  Watcher stopped")
}

func watchDir(dir string, langs []string, done chan bool) {
	fileTimes := make(map[string]time.Time)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".bak") || strings.Contains(path, "/tmp/") {
			return nil
		}
		fileTimes[path] = info.ModTime()
		return nil
	})

	for {
		time.Sleep(500 * time.Millisecond)

		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".bak") || strings.Contains(path, "/tmp/") {
				return nil
			}

			lastTime, exists := fileTimes[path]
			if !exists || info.ModTime().After(lastTime) {
				fmt.Printf(" 🔄  Changed: %s\n", path)
				fileTimes[path] = info.ModTime()
				runBuild(langs, false)
			}
			return nil
		})
	}
}
