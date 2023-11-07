package historics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers"
	"supVOD/app/models"
	"supVOD/app/services/historics"
	"supVOD/app/services/movies"
	"supVOD/app/services/users"
)

func getUserAndMovie(h models.Historics) (models.User, models.Movie, error) {
	user, err := users.GetUserByID(h.UserID)
	if err != nil {
		return models.User{}, models.Movie{}, fmt.Errorf(handlers.ErrorMovieNotFound)
	}

	movie, err := movies.GetMovieByID(h.MovieID)
	if err != nil {
		return models.User{}, models.Movie{}, fmt.Errorf(handlers.ErrorMovieNotFound)
	}

	return *user, *movie, nil
}

func CreateHistorics(c *gin.Context) {
	var requestHistorics models.Historics
	err := c.BindJSON(&requestHistorics)
	if err != nil || !requestHistorics.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, requestHistorics.RequiredFieldsString()))
		return
	}

	user, movie, err := getUserAndMovie(requestHistorics)
	if err != nil {
		handlers.ErrorResponse(c, 400, err.Error())
		return
	}

	err = historics.Create(&requestHistorics)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, requestHistorics.ToRichMap(user, movie))
}

func FindHistorics(c *gin.Context) {
	ID := c.Query("user_id")
	historicsList, err := historics.Find(ID)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	resp := make([]map[string]interface{}, len(historicsList))
	for i, historic := range historicsList {
		resp[i] = historic.ToMap()
	}

	handlers.SuccessResponse(c, 201, map[string]interface{}{
		"historics": resp,
	})
}

func GetHistoricsByID(c *gin.Context) {
	ID := c.Param("id")
	requestHistorics, err := historics.GetById(ID)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	user, movie, err := getUserAndMovie(*requestHistorics)
	if err != nil {
		handlers.ErrorResponse(c, 400, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, requestHistorics.ToRichMap(user, movie))
}
