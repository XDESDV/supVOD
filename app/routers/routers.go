package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"supVOD/app/routers/kinds"
	"supVOD/app/routers/movies"
	"supVOD/app/routers/users"
	"time"
)

type Server struct {
	server *gin.Engine
}

func (s *Server) Init() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(cors.Default())

	users.InitRouter(r)
	movies.InitRouter(r)
	kinds.InitRouter(r)

	s.server = r
}

func (s *Server) Run() error {
	fmt.Println(fmt.Sprintf("Server running on port %s", os.Getenv("APP_PORT")))
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:        s.server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	return err
}
