package handlers

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func SuccessResponse(c *gin.Context, code int, data map[string]any) {
	c.JSON(code, data)
}

const (
	ErrorBadBody       string = "Invalid request body %s"
	ErrorUserNotFound  string = "User not found"
	ErrorMovieNotFound string = "User not found"
	ErrorKindNotFound  string = "Kind not found %s"
	ErrorInternal      string = "Internal server error"
)
