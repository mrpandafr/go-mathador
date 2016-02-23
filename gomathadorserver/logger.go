package gomathadorserver

import (
	"log"
	"net/http"
	"time"

	"github.com/mrpandafr/go-mathador/gomathadorredis"
)

//Logger : output server request
func Logger(inner http.Handler, name string, c gomathadorredis.Channel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		c.Publish("test")
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
