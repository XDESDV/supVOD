package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// initialiser le routeur
func InitialiseRouter() *gin.Engine {
	log.Println("Initialise Router")
	//router := gin.Default()
	//router.GET("/users")
	return gin.Default()
}
