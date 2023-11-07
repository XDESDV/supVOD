package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitRouter
func InitialiseRouter() *gin.Engine {
	log.Println("Initialise Router")

	return gin.Default()
}
