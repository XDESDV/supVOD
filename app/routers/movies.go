package routers

import (
	"supVOD/app/handlers"

	"github.com/gin-gonic/gin"
)

// InitialiseRouter is a function.
func InitialiseMoviesRoute(g *gin.Engine) {
	g.GET("/movies/:id", handlers.GetMovieByID)
	g.GET("/movies", handlers.GetAllMovies)
	g.POST("movies", handlers.CreateMovie)
	g.POST("movies/:id", handlers.UpdateMovie)
}
