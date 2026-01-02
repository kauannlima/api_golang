package main

import (
	"github.com/kauannlima/api_golang/config"
	"github.com/kauannlima/api_golang/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	//Initialize the database
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize the router
	router.Initialize()
}
