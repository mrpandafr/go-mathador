package gomathadorserver

import "net/http"

//Route : struct for REST and Web route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes : list of Route
type Routes []Route
