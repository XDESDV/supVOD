package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"supVOD/app/models"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var m models.Movie

	m.Title = c.PostForm("title")
	m.Description = c.PostForm("description")

	if duration, err := strconv.Atoi(c.PostForm("duration")); err == nil {
		m.Duration = duration
	}

	var kinds []models.Kinds
	if err := json.Unmarshal([]byte(c.PostForm("kinds")), &kinds); err == nil {
		m.Kinds = kinds
	}

	err := services.CreateMovie(m)
	if err != nil {
		//Gestion des erreurs
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, nil)
}

func UpdateMovie(c *gin.Context) {
	var m models.Movie

	// log.Println("TEST UPDATE")

	// m.Title = c.PostForm("title")
	// m.Description = c.PostForm("description")

	// if duration, err := strconv.Atoi(c.PostForm("duration")); err == nil {
	// 	m.Duration = duration
	// }

	// var kinds []models.Kinds
	// if err := json.Unmarshal([]byte(c.PostForm("kinds")), &kinds); err == nil {
	// 	m.Kinds = kinds
	// }

	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	log.Println("TEST UPDATE 2")

	err := services.UpdateMovie(m)
	if err != nil {
		//Gestion des erreurs
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, nil)
}

func GetMovieByID(c *gin.Context) {
	id := c.Params.ByName("id")

	movie, err := services.GetMovieByID(id)
	if err != nil {
		//Gestion des erreurs
		c.JSON(http.StatusInternalServerError, err)
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, movie)
}

func GetAllMovies(c *gin.Context) {
	movie, err := services.GetMovies()
	if err != nil {
		//Gestion des erreurs
		c.JSON(http.StatusInternalServerError, err)
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	log.Println("everything good!")
	c.JSON(http.StatusOK, movie)
}
