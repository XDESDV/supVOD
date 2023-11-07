package movies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"supVOD/app/handlers"
	"supVOD/app/models"
	"supVOD/app/services/kinds"
	"supVOD/app/services/movies"
)

func checkKinds(k *[]models.Kind) error {
	if (k != nil) && (len(*k) > 0) {
		kindList, err := kinds.List()
		if err != nil {
			return fmt.Errorf(handlers.ErrorInternal)
		}

	KindList:
		for i := range *k {
			for _, kind := range kindList {
				if kind.Name == (*k)[i].Name ||
					kind.ID == (*k)[i].ID {
					(*k)[i] = models.Kind{ID: kind.ID, Name: kind.Name}
					continue KindList
				}
			}
			return fmt.Errorf(fmt.Sprintf(handlers.ErrorKindNotFound, (*k)[i].Name))
		}
	}

	return nil
}

func CreateMovie(c *gin.Context) {
	var requestMovie models.Movie
	err := c.BindJSON(&requestMovie)
	if err != nil || !requestMovie.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, requestMovie.RequiredFieldsString()))
		return
	}

	err = checkKinds(&requestMovie.Kinds)
	if err != nil {
		handlers.ErrorResponse(c, 400, err.Error())
		return
	}

	err = movies.Create(&requestMovie)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, requestMovie.ToMap())
}

func ListMovie(c *gin.Context) {
	var page int
	var limit int
	var err error

	qPage := c.Query("page")
	qLimit := c.Query("limit")
	qSearch := c.Query("search")
	qKinds := c.QueryArray("kinds")

	if qPage == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(qPage)
		if err != nil {
			page = 1
		}
	}

	if qLimit == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			limit = 10
		}
	}

	var kindsFromQuery []models.Kind
	for _, kind := range qKinds {
		kindsFromQuery = append(kindsFromQuery, models.Kind{
			Name: kind,
		})
	}

	movieList, err := movies.List(
		models.PaginationParams{
			Page:     page,
			PageSize: limit},
		models.SearchParams{
			Search: qSearch,
			Kinds:  kindsFromQuery,
		})
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	resp := make([]map[string]interface{}, len(movieList.Movies))

	for i, movie := range movieList.Movies {
		resp[i] = movie.ToMap()
	}

	handlers.SuccessResponse(c, 200, map[string]interface{}{
		"movies":   resp,
		"page":     page,
		"limit":    limit,
		"total":    movieList.TotalCount,
		"max_page": movieList.MaxPage,
	})
}

func GetMovieByID(c *gin.Context) {
	ID := c.Param("id")

	movieRequested, err := movies.GetMovieByID(ID)
	if err != nil {
		handlers.ErrorResponse(c, 404, handlers.ErrorMovieNotFound)
		return
	}

	handlers.SuccessResponse(c, 200, movieRequested.ToMap())
}

func UpdateMovie(c *gin.Context) {
	var newMovie models.Movie
	err := c.BindJSON(&newMovie)
	if err != nil || !newMovie.Validate() {
		fmt.Println(err)
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, newMovie.RequiredFieldsString()))
		return
	}

	err = checkKinds(&newMovie.Kinds)
	if err != nil {
		handlers.ErrorResponse(c, 400, err.Error())
		return
	}

	_, err = movies.GetMovieByID(newMovie.ID)
	if err != nil {
		handlers.ErrorResponse(c, 404, handlers.ErrorUserNotFound)
		return
	}

	err = movies.UpdateMovie(&newMovie)
	if err != nil {
		handlers.ErrorResponse(c, 500, handlers.ErrorInternal)
		return
	}

	handlers.SuccessResponse(c, 201, newMovie.ToMap())
}
