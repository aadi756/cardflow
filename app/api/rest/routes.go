package rest

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
	"github.com/bnkamalesh/webgo/middleware"
)

// Routes returns a list of handlers defined for webGo
func GetRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:     "helloworld",                   // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                 // request type
			Pattern:  "/",                            // Pattern for the route
			Handlers: []http.HandlerFunc{helloWorld}, // route handler
		},
		&webgo.Route{
			Name:     "helloworld",                                      // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                    // request type
			Pattern:  "/api/:param",                                     // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), helloWorld}, // route handler
		},
	}
}
