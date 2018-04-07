package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	//Handle with the action form submit
	http.HandleFunc("/markdown", GenerateMarkdown)
	//Handle with the root page
	http.Handle("/", http.FileServer(http.Dir("public")))

	//Initialize server

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
