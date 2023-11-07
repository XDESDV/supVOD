package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitialiseRouter initialisation des routes du web service
func InitialiseRouter() *gin.Engine {
	log.Println("Initialise router")
	router := gin.Default()
	return router
}
