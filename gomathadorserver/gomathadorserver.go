package gomathadorserver

import (
	"github.com/gorilla/mux"
	"github.com/mrpandafr/go-mathador/gomathadorredis"
)

type server struct {
	Channel gomathadorredis.Channel
	router  *mux.Router
}

// Server : commun web and rest interface
type Server interface {
}

// NewServer : init and return a web and rest server
func NewServer(c gomathadorredis.Channel) Server {
	s := &server{
		Channel: c,
		router:  NewRoutes(),
	}
	return s
}
