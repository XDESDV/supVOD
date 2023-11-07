package kinds

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"supVOD/app/handlers"
	"supVOD/app/models"
	"supVOD/app/services/kinds"
)

func CreateKind(c *gin.Context) {
	var requestKind models.Kind
	err := c.BindJSON(&requestKind)
	if err != nil || !requestKind.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, requestKind.RequiredFieldsString()))
		return
	}

	requestKind.Name = strings.ToLower(requestKind.Name)
	err = kinds.Create(&requestKind)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, requestKind.ToMap())
}

func UpdateKind(c *gin.Context) {
	var requestKind models.Kind
	err := c.BindJSON(&requestKind)
	if err != nil || !requestKind.Validate() {
		handlers.ErrorResponse(c, 400, fmt.Sprintf(handlers.ErrorBadBody, requestKind.RequiredFieldsString()))
		return
	}

	requestKind.Name = strings.ToLower(requestKind.Name)
	err = kinds.Update(&requestKind)
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	handlers.SuccessResponse(c, 201, requestKind.ToMap())
}

func ListKind(c *gin.Context) {
	kindList, err := kinds.List()
	if err != nil {
		handlers.ErrorResponse(c, 500, err.Error())
		return
	}

	resp := make([]map[string]interface{}, len(kindList))
	for i, kind := range kindList {
		resp[i] = kind.ToMap()
	}

	handlers.SuccessResponse(c, 200, map[string]interface{}{
		"kinds": resp,
	})
}
