package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	//Collection handlers
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	//Post single request list
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

//HomeHandler controller is a function to handle with root request
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "home")
}

//PostsIndexHandler controller is a function to handle with the list of all posts
func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

//PostsCreateHandler controller is a function to handle with the save action
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

//PostShowHandler controller is a function to handle with the save action
func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

//PostUpdateHandler controller is a function to handle with the save action
func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

//PostDeleteHandler controller is a function to handle with the save action
func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

//PostEditHandler controller is a function to handle with the save action
func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
