package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"supVOD/app/models"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

// Find GetbyId Create Update

func CreateMovie(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var m models.Movie
	err := json.NewDecoder(req.Body).Decode(&m)
	messagesTypes := &models.MessageTypes{
		Created:      "movie.create.done",
		BadRequest:   "movie.create.badrequest",
		Unauthorized: "movie.create.failed"}
	if err == nil {
		if m.Title != "" {
			if m.Description != "" {
				if m.Duration > 0 {
					if len(m.Kinds) > 0 {
						if err := services.CreateMovie(m); err == nil {
							SuccessResponse(w, http.StatusCreated, messagesTypes.Created, "movie created")
						} else {
							ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
						}
					} else {
						err = errors.New("empty kind(s)")
						ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
					}
				} else {
					err = errors.New("empty or invalid duration")
					ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
				}
			} else {
				err = errors.New("empty description")
				ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
			}
		} else {
			err = errors.New("empty title")
			ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
		}
	} else {
		err = errors.New("error decode")
		ErrorResponse(w, http.StatusUnauthorized, messagesTypes, err)
	}
}

func UpdateMovie(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var m models.Movie
	messagesTypes := &models.MessageTypes{
		Created:      "movie.update.done",
		BadRequest:   "movie.update.badrequest",
		Unauthorized: "movie.update.failed"}
	err := json.NewDecoder(req.Body).Decode(&m)
	if err == nil {
		m.ID = c.Param("id")
		if err := services.UpdateMovie(m); err == nil {
			SuccessResponse(w, http.StatusCreated, messagesTypes.Created, "movie updated")
		} else {
			ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
		}
	} else {
		err = errors.New("error decode")
		ErrorResponse(w, http.StatusUnauthorized, messagesTypes, err)
	}
}

func GetMoviebyId(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	messagesTypes := &models.MessageTypes{
		Created:      "movie.getbyid.done",
		BadRequest:   "movie.getbyid.badrequest",
		Unauthorized: "movie.getbyid.failed"}
	if m, err := services.GetMovieByID(c.Param("id")); err == nil {
		var wsr models.WSResponse
		wsr.Meta.ObjectName = "movie"
		wsr.Meta.TotalCount = 1
		wsr.Meta.Offset = 1
		wsr.Meta.Count = 1
		wsr.Data = m
		ReturnResponse(w, http.StatusFound, wsr)
	} else {
		ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
	}

}

func FindMovie(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var qm models.Query_Movie
	messagesTypes := &models.MessageTypes{
		Created:      "movie.getbyid.done",
		BadRequest:   "movie.getbyid.badrequest",
		Unauthorized: "movie.getbyid.failed"}
	for key, values := range req.URL.Query() {
		switch key {
		case "id":
			qm.IDs = values
			break
		case "title":
			qm.Titles = values
			break
		case "kind":
			qm.Kinds = values
			break
		default:
			break
		}
	}
	if m, err := services.FindMovie(qm); err == nil {
		var wsr models.WSResponse
		wsr.Meta.ObjectName = "movie"
		wsr.Data = m
		ReturnResponse(w, http.StatusFound, wsr)
	} else {
		ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
	}
}
