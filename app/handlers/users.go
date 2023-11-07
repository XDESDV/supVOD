package handlers

import (
	"net/http"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {

}

func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := services.GetByID(id)
	if err != nil {
		//Gestion des erreurs
		c.JSON(http.StatusInternalServerError, err)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, user)
}
