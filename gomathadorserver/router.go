package gomathadorserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRoutes : open a web service initialise routes
func NewRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
