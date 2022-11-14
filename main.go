package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/gl":  "https://www.google.com",
		"/git": "https://www.github.com/vyas-git",
	}

}
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
