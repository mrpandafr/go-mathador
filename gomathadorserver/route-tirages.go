package gomathadorserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func tiragesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tirages Index!")
}

func tiragesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTirage := vars["idTirage"]
	fmt.Fprintln(w, "Tirage show:", idTirage)
}
