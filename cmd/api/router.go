package main

import (
	"net/http"
	"supVOD/app/handlers"
	"supVOD/app/routers"

	"github.com/gin-gonic/gin"
)

func health(w http.ResponseWriter, req *http.Request) {

}

func initialiseRouter() *gin.Engine {
	r := routers.InitialiseRouter()
	r.GET("/movie", handlers.FindMovie)
	r.POST("/movie", handlers.CreateMovie)
	r.POST("/movie/:id", handlers.UpdateMovie)
	r.GET("/movie/:id", handlers.GetMoviebyId)
	return r
}
