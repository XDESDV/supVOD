package routers

import (
	"supVOD/app/handlers"

	"github.com/gin-gonic/gin"
)

func InitialiseMoviesRoute(g *gin.Engine) {

	g.GET("/movies/:id", handlers.GetMovieByID)

}
