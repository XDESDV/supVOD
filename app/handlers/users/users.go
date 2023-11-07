package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers"
	"supVOD/app/models"
	"supVOD/app/services/users"
)

func CreateUser(c *gin.Context) {
	var requestUser models.User
	err := c.BindJSON(&requestUser)
	if err != nil || !requestUser.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, requestUser.RequiredFieldsString()))
		return
	}

	err = users.CreateUser(requestUser)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, map[string]interface{}{
		"id":    requestUser.ID,
		"email": requestUser.Email,
	})
}

func GetUserByID(c *gin.Context) {
	ID := c.Param("id")

	user, err := users.GetUserByID(ID)
	if err != nil {
		handlers.ErrorResponse(c, 404, handlers.ErrorUserNotFound)
		return
	}

	handlers.SuccessResponse(c, 200, user.ToMap())
}

func UpdateUser(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil || !newUser.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, newUser.RequiredFieldsString()))
		return
	}

	_, err = users.GetUserByID(newUser.ID)
	if err != nil {
		handlers.ErrorResponse(c, 404, handlers.ErrorUserNotFound)
		return
	}

	err = users.UpdateUser(newUser)
	if err != nil {
		handlers.ErrorResponse(c, 500, handlers.ErrorInternal)
		return
	}

	handlers.SuccessResponse(c, 201, newUser.ToMap())
}
