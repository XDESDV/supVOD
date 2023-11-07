package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitialiseRouter() *gin.Engine {
	log.Println("Initialising router")
	router := gin.Default()
	return router

	// return gin.Default()
}
