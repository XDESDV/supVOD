package handlers

import (
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateUser() {

}

func GetMovieByID(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := services.GetMovieByID(id)
	if err != nil {
		// gestion de l'erreur
		c.JSON(http.StatusInternalServerError, err)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, nil)
	}
	c.JSON(http.StatusOK, user)

}
