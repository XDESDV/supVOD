package movies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers/movies"
)

func InitRouter(r *gin.Engine) {
	fmt.Println("Initializing movies router...")
	v1 := r.Group("/v1/movies")
	{
		v1.POST("/", movies.CreateMovie)
		v1.GET("/", movies.ListMovie)
		v1.GET("/:id", movies.GetMovieByID)
		v1.PUT("/", movies.UpdateMovie)
	}
}
