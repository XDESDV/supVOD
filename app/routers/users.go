package routers

import (
	"supVOD/app/handlers"

	"github.com/gin-gonic/gin"
)

func InitialiseUsersRoute(g *gin.Engine) {

	g.GET("/users/:id", handlers.GetByID)

}
