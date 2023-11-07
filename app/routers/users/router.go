package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"supVOD/app/handlers/users"
)

func InitRouter(r *gin.Engine) {
	fmt.Println("Initializing users router...")
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", users.CreateUser)
		v1.PUT("/", users.UpdateUser)
		v1.GET("/:id", users.GetUserByID)
	}
}
