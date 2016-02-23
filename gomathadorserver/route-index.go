package gomathadorserver

import (
	"fmt"
	"html"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour sur le serveur go-mathador (Handle with care) %q", html.EscapeString(r.URL.Path))
}
