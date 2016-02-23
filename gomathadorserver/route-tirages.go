package gomathadorserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpandafr/go-mathador/model"
)

func tirageIndex(w http.ResponseWriter, r *http.Request) {
	tirages := model.Tirages{
		model.Tirage{De1: 1, De2: 2, De3: 3, De4: 4, De5: 5},
		model.Tirage{De1: 1, De2: 2, De3: 3, De4: 4, De5: 6},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tirages); err != nil {
		panic(err)
	}
}

func tirageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTirage := vars["idTirage"]
	fmt.Fprintln(w, "Tirage show:", idTirage)
}
