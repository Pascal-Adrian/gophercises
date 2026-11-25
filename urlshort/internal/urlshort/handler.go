package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.yaml.in/yaml/v4"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		val, ok := pathsToUrls[url]

		if ok {
			fmt.Printf("Redirect: %s -> %s\n", url, val)
			http.Redirect(w, r, val, http.StatusFound)
		} else {
			fmt.Printf("Get: %s\n", url)
			fallback.ServeHTTP(w, r)
		}

	}
}

// StructuredHandler is the signature shared by handlers that
// build their path-to-URL mapping from structured data (such
// as YAML or JSON). It parses in and returns an
// http.HandlerFunc (which also implements http.Handler) that
// will attempt to map any paths to their corresponding URL.
// If the path is not present in the data, then the fallback
// http.Handler will be called instead.
//
// The data is expected to be a list of entries, each with a
// "path" and a "url" field.
//
// The only errors that can be returned are all related to
// having invalid data for the given format.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type StructuredHandler func(in []byte, fallback http.Handler) (http.HandlerFunc, error)

// YAMLHandler is a [StructuredHandler] for YAML data.
//
// YAML is expected to be in the format:
//
//	---
//	- path: /some-path
//	  url: https://www.some-url.com/demo
func YAMLHandler(in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parse(in, yaml.Unmarshal)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler is a [StructuredHandler] for JSON data.
//
// JSON is expected to be in the format:
//
//	[
//	  {"path": "/some-path", "url": "https://www.some-url.com/demo"}
//	]
func JSONHandler(in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parse(in, json.Unmarshal)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathsToUrls, fallback), nil
}
