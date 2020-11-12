package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		decodeToken := true
		if decodeToken != true {
			fmt.Fprint(res, "\n\nINVALID USER\n\n")
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Fprint(res, "\n\nMIDDLEWARE METHOD\n\n")
		next(res, req)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "\n\nTHIS IS INDEX\n\n")
}

func show(res http.ResponseWriter, req *http.Request) {
	id := req.URL
	fmt.Println("URL IS :", id)
	fmt.Fprint(res, "\n\nTHIS IS SHOW\n\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", authenticate(show))
	r.HandleFunc("/users", index)
	http.ListenAndServe(":8000", r)
}
