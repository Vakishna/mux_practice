package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
