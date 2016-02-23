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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index,
	},
	Route{
		"TirageIndex",
		"GET",
		"/tirages",
		tirageIndex,
	},
	Route{
		"TirageShow",
		"GET",
		"/tirages/{tirageId}",
		tirageShow,
	},
}
