package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"supVOD/app/models"
	"supVOD/app/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var u models.User
	err := json.NewDecoder(req.Body).Decode(&u)
	messagesTypes := &models.MessageTypes{
		Created:      "user.create.done",
		BadRequest:   "user.create.badrequest",
		Unauthorized: "user.create.failed"}
	if err == nil {
		if u.Email != "" {
			if u.UserPassword != "" {
				if err := services.CreateUser(u); err == nil {
					SuccessResponse(w, http.StatusCreated, messagesTypes.Created, "user created")
				} else {
					ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
				}
			} else {
				err = errors.New("empty password")
				ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
			}
		} else {
			err = errors.New("empty email")
			ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
		}
	} else {
		err = errors.New("error decode")
		ErrorResponse(w, http.StatusUnauthorized, messagesTypes, err)
	}
}

func UpdateUser(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var u models.User
	messagesTypes := &models.MessageTypes{
		Created:      "user.update.done",
		BadRequest:   "user.update.badrequest",
		Unauthorized: "user.update.failed"}
	err := json.NewDecoder(req.Body).Decode(&u)
	if err == nil {
		u.ID = c.Param("id")
		if err := services.UpdateUser(u); err == nil {
			SuccessResponse(w, http.StatusCreated, messagesTypes.Created, "user updated")
		} else {
			ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
		}
	} else {
		err = errors.New("error decode")
		ErrorResponse(w, http.StatusUnauthorized, messagesTypes, err)
	}
}
