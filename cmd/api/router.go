package main

import (
	"net/http"
	"supVOD/app/routers"

	"github.com/gin-gonic/gin"
)

func health(w http.ResponseWriter, req *http.Request) {

}

func initialiseRouter() *gin.Engine {
	r := routers.InitialiseRouter()

	return r
}
