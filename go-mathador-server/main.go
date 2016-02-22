package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// InitREST open a web service initialise routes
func InitREST() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/tirages", tiragesIndex)
	router.HandleFunc("/tirages/{idTirage}", tiragesShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func tiragesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tirages Index!")
}

func tiragesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTirage := vars["idTirage"]
	fmt.Fprintln(w, "Tirage show:", idTirage)
}

// TODO: Make a REST server ASAP
func main() {
	fmt.Print("dummy server")
	InitREST()
}
