package routers

import (
	"supVOD/app/handlers"

	"github.com/gin-gonic/gin"
)

// InitialiseRouter is a function.
func InitialiseUsersRoute(g *gin.Engine) {
	g.GET("/users/:id", handlers.GetByID)
	// g.POST("/users", handlers.CreateUser)
	// g.POST("/users/id", handlers.UpdateUser)
}
