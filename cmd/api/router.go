package main

import (
	"supVOD/app/routers"

	"github.com/gin-gonic/gin"
)

func initialiseRouter() *gin.Engine {

	r := routers.InitialiseRouter()

	routers.InitialiseUsersRoute(r)

	routers.InitialiseMoviesRoute(r)

	return r
}
