package handlers

import (
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	movie, err := services.GetMovieByID(id)
	if err != nil {
		// gestion de l'erreur
		c.JSON(http.StatusInternalServerError, err)
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, nil)
	}
	c.JSON(http.StatusOK, movie)

}
