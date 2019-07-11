package main

import (
	"os"
	"time"

	"github.com/bnkamalesh/webgo/middleware"

	"github.com/aadi756/cardflow/app/api/rest"
	"github.com/aadi756/cardflow/app/pkg/platform/db"
	"github.com/bnkamalesh/webgo"
)

func main() {
	cfg := webgo.Config{
		Host:         "",
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	m := db.Config{
		Host : []string{"127.0.0.1:27017"},
		Port : "27017",
		Name : "credit",
		Username : "",
		Password : "",
		Direct : true,
		Timeout : time.Second * 10,
	}
	// m, err := configs.DataStore()
	// if err != nil {
	// 	os.Exit(1)
	// 	return
	// }


	//dbhandler, err := db.New(m)
	//if err != nil {
	//	os.Exit(1)
	//	return
	//}

	err := db.Init(m)
	if err != nil {
		os.Exit(1)
		return
	}
	router := webgo.NewRouter(&cfg, rest.GetRoutes())
	router.Use(middleware.AccessLog)
	router.Start()
}
