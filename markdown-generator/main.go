package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	//Handle with the action form submit
	http.HandleFunc("/markdown", GenerateMarkdown)
	//Handle with the root page
	http.Handle("/", http.FileServer(http.Dir("public")))

	//Initialize server
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
