// handlers/movie.go
package handlers

import (
	"log"
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
}

func GetMovieByID(c *gin.Context) {
	id := c.Params.ByName("id")

	movie, err := services.GetMovieByID(id)

	// Gestion de l'erreur 500
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	// Gestion de l'erreur 404
	if movie == nil {
		c.JSON(http.StatusNotFound, nil)
	}
	log.Println(movie)
	// Gestion OK
	c.JSON(http.StatusOK, movie)
}
