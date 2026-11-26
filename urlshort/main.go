package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Pascal-Adrian/gophercises/urlshort/internal/database"
	"github.com/Pascal-Adrian/gophercises/urlshort/internal/urlshort"

	dotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func run() (int, error) {
	pathname := flag.String("filepath", "paths.yaml", "filepath to the YAML/JSON file with the paths")

	flag.Parse()

	err := dotenv.Load(".env")

	if err != nil {
		fmt.Println("failed to load env file")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		return 2, errors.New("no DB connection string provided")
	}

	fileType := strings.ToLower(filepath.Ext(*pathname))

	var handlerFunc urlshort.StructuredHandler

	switch fileType {
	case ".json":
		handlerFunc = urlshort.JSONHandler
	case ".yml", ".yaml":
		handlerFunc = urlshort.YAMLHandler
	default:
		return 2, errors.New("unsupported file type")
	}

	data, err := os.ReadFile(*pathname)

	if err != nil {
		return 1, err
	}

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	handler, err := handlerFunc(data, mapHandler)

	if err != nil {
		return 1, err
	}

	connection, err := sql.Open("sqlite3", dbUrl)

	if err != nil {
		return 1, err
	}

	defer connection.Close()

	dbHandler := urlshort.DBHandler(database.New(connection), handler)

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", dbHandler)

	return 1, err
}

func main() {
	code, err := run()

	fmt.Fprintln(os.Stderr, err)
	os.Exit(code)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
