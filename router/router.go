package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize the Gin router
	router := gin.Default()

	//Initialize routes
	InitializeRoutes(router)

	//Run the router
	router.Run(":8080")
}
