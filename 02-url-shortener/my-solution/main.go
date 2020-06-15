package main

import (
	"fmt"
	"gophercises/02-url-shortener/my-solution/urlshort"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlString := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlString), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")

	// m := make(map[interface{}]interface{})
	// err = yaml.Unmarshal([]byte(yamlString), &m)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("WOOOWHOO", m)

	http.ListenAndServe(":8080", yamlHandler)
	// http.ListenAndServe(":8080", mapHandler)

	_ = yamlHandler
	_ = mapHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
