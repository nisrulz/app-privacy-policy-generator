package main

import (
	"fmt"
	"net/http"
	"os"
)

func runServe(langs []string, clean bool, port string) {
	runBuild(langs, clean)
	fmt.Printf(" 🌐  Serving public/ on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, http.FileServer(http.Dir("public"))); err != nil {
		fmt.Fprintf(os.Stderr, " ❌  %v\n", err)
	}
}
