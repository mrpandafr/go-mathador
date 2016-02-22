package gomathadorserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes : open a web service initialise routes
func Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/tirages", tiragesIndex)
	router.HandleFunc("/tirages/{idTirage}", tiragesShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
