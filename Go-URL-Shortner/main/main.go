package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/aditya-nagare/Portfolio/go-url-shortner"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/ggl": "https://google.com",
		"/yt":  "https://youtube.com",
		"/fb":  "https://facebook.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /wiki
  url: https://wikipedia.org
- path: /med
  url: https://medium.com
- path: /9
  url: https://9gag.com
- path: /urlshort
  url: https://github.com/aditya-nagare/Portfolio/Go-URL-Shortner
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
