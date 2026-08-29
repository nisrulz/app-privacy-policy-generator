package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")

		next.ServeHTTP(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

func runServe(langs []string, clean bool, port string) {
	runBuild(langs, clean)
	fmt.Printf(" 🌐  Serving public/ on http://localhost:%s\n", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: gzipMiddleware(http.FileServer(http.Dir("public"))),
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n 🛑  Server stopped")
		server.Shutdown(context.Background())
		os.Exit(0)
	}()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		if isPortInUse(err) {
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, " ------------------------------------------------")
			fmt.Fprintf(os.Stderr, " ❌  Port %s is already in use.\n\n", port)
			fmt.Fprintln(os.Stderr, " 👨🏻‍💻  Run with a different port: make serve PORT=\"8080\"")
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, " ❌  %v\n", err)
		os.Exit(1)
	}
}

func isPortInUse(err error) bool {
	return strings.Contains(err.Error(), "address already in use")
}
