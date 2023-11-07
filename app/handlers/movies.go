package handlers

import (
	"log"
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateMovie() {

}

func GetByIDMovie(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := services.GetByID(id)

	// Gestion de l'erreur 500
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	// Gestion de l'erreur 404
	if user == nil {
		c.JSON(http.StatusNotFound, nil)
	}
	log.Println(user)

	// Gestion OK
	c.JSON(http.StatusOK, user)
}
