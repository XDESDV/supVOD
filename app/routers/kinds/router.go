package kinds

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers/kinds"
)

func InitRouter(r *gin.Engine) {
	fmt.Println("Initializing kinds router...")
	v1 := r.Group("/v1/kinds")
	{
		v1.POST("/", kinds.CreateKind)
		v1.GET("/", kinds.ListKind)
		v1.PUT("/", kinds.UpdateKind)
	}
}
