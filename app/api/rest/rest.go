package rest

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	wctx := webgo.Context(r)
	webgo.R200(
		w,
		wctx.Params, // URI parameters
	)
}
