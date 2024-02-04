package main

import (
	"fmt"
	"net/http"

	urlshortener "github.com/bishalbera/url-shortener"
)

func main() {
	mux:= defaultMux()
	pathsToUrls:= map[string]string {
		"/google": "https://google.com",
		"/youtube": "https://youtube.com",
	}

	mapHandler:= urlshortener.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8085")
	http.ListenAndServe(":8085", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", love)
	return mux
}

func love(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Love You 3000 :)")
}