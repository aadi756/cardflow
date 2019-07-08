package main

import (
	"time"

	"github.com/bnkamalesh/webgo/middleware"

	"github.com/aadi756/cardflow/app/api/rest"
	"github.com/bnkamalesh/webgo"
	_ "rsc.io/quote"
)

func main() {
	cfg := webgo.Config{
		Host:         "",
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	router := webgo.NewRouter(&cfg, rest.GetRoutes())
	router.Use(middleware.AccessLog)
	router.Start()
}
