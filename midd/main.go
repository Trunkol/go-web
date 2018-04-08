package main

import "github.com/codegangsta/negroni"

func main() {
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public"))
	)
	n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)  {
	log.println("Add log in  the way there...")

	if r.URL.Query().Get("password") == "secret123"{
		next(rw, r)
	} else{
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging the way back")
}