package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	urlshort "github.com/aditya-nagare/Portfolio/go-url-shortner"
)

func main() {
	yamlFilename := flag.String("yaml", "links.yaml", "Yaml file name with redirection URLs")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/ggl": "https://google.com",
		"/yt":  "https://youtube.com",
		"/fb":  "https://facebook.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// 	yaml := `
	// - path: /wiki
	//   url: https://wikipedia.org
	// - path: /med
	//   url: https://medium.com
	// - path: /9
	//   url: https://9gag.com
	// - path: /urlshort
	//   url: https://github.com/aditya-nagare/Portfolio/Go-URL-Shortner
	// `

	yamlFile, err := ioutil.ReadFile(*yamlFilename)
	if err != nil {
		fmt.Printf("Failed to Open the YAML File:%s", *yamlFilename)
		os.Exit(1)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
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
