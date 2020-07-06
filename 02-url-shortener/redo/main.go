package main

import (
	"fmt"
	"net/http"
	"sandbox/gophercises/02-url-shortener/redo/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8000")

	// http.ListenAndServe(":8000", mapHandler)
	http.ListenAndServe(":8000", yamlHandler)

	_ = yamlHandler
	_ = mapHandler

}

func defaultMux() *http.ServeMux {
	hf := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! you are at %s", r.URL.Path)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(hf))
	return mux
}
