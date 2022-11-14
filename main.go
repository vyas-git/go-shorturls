package main

import (
	"fmt"
	"net/http"

	"github.com/vyas-git/go-shorturls/urlshort"
)

func main() {
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/gl":  "https://www.google.com",
		"/git": "https://www.github.com/vyas-git",
	}
	maphandler := urlshort.MapHandler(pathToUrls, mux)
	fmt.Println("Server starting on 8080")
	http.ListenAndServe(":8080", maphandler)
}
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
