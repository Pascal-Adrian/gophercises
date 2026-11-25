package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Pascal-Adrian/gophercises/urlshort/internal/urlshort"
)

func exitWithError(message string, code int) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}

func main() {
	pathname := flag.String("filepath", "paths.yaml", "filepath to the YAML/JSON file with the paths")

	flag.Parse()

	fileType := strings.ToLower(filepath.Ext(*pathname))

	var handlerFunc urlshort.StructuredHandler

	switch fileType {
	case ".json":
		handlerFunc = urlshort.JSONHandler
	case ".yml", ".yaml":
		handlerFunc = urlshort.YAMLHandler
	default:
		exitWithError("Unsupported file type", 2)
	}

	data, err := os.ReadFile(*pathname)

	if err != nil {
		exitWithError(err.Error(), 1)
	}

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	handler, err := handlerFunc(data, mapHandler)

	if err != nil {
		exitWithError(err.Error(), 1)
	}

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", handler)

	if err != nil {
		exitWithError(err.Error(), 1)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
