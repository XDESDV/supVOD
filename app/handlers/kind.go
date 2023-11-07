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

func CreateKind(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var k models.Kind
	err := json.NewDecoder(req.Body).Decode(&k)
	messagesTypes := &models.MessageTypes{
		Created:      "kind.create.done",
		BadRequest:   "kind.create.badrequest",
		Unauthorized: "kind.create.failed"}
	if err == nil {
		if k.Name != "" {
			if err := services.CreateKind(k); err == nil {
				SuccessResponse(w, http.StatusCreated, messagesTypes.Created, "kind created")
			} else {
				ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
			}
		} else {
			err = errors.New("empty name")
			ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
		}
	} else {
		err = errors.New("error decode")
		ErrorResponse(w, http.StatusUnauthorized, messagesTypes, err)
	}
}

func GetKindbyId(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	messagesTypes := &models.MessageTypes{
		Created:      "kind.getbyid.done",
		BadRequest:   "kind.getbyid.badrequest",
		Unauthorized: "kind.getbyid.failed"}
	if m, err := services.GetKindByID(c.Param("id")); err == nil {
		var wsr models.WSResponse
		wsr.Meta.ObjectName = "kind"
		wsr.Meta.TotalCount = 1
		wsr.Meta.Offset = 1
		wsr.Meta.Count = 1
		wsr.Data = m
		ReturnResponse(w, http.StatusFound, wsr)
	} else {
		ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
	}

}

func FindKind(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	var qk models.Query_Kind
	messagesTypes := &models.MessageTypes{
		Created:      "kind.getbyid.done",
		BadRequest:   "kind.getbyid.badrequest",
		Unauthorized: "kind.getbyid.failed"}
	for key, values := range req.URL.Query() {
		switch key {
		case "id":
			qk.IDs = values
			break
		case "name":
			qk.Names = values
			break
		default:
			break
		}
	}
	if m, err := services.FindKind(qk); err == nil {
		var wsr models.WSResponse
		wsr.Meta.ObjectName = "kind"
		wsr.Data = m
		ReturnResponse(w, http.StatusFound, wsr)
	} else {
		ErrorResponse(w, http.StatusBadRequest, messagesTypes, err)
	}
}
