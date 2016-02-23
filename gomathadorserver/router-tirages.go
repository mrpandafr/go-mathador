package gomathadorserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func tirageIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tirages Index!")
}

func tirageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTirage := vars["idTirage"]
	fmt.Fprintln(w, "Tirage show:", idTirage)
}
