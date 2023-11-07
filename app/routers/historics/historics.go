package historics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers/historics"
)

func InitRouter(r *gin.Engine) {
	fmt.Println("Initializing historics router...")
	v1 := r.Group("/v1/historics")
	{
		v1.POST("/", historics.CreateHistorics)
		v1.GET("/", historics.FindHistorics)
		v1.GET("/:id", historics.GetHistoricsByID)
	}
}
