package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Host   string
	router *gin.Engine
}

func New(host string) *Server {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.POST("/regestratin")
		userGroup.POST("/auth")
		userGroup.GET("/")
	}

	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/all-books")
		bookGroup.GET("/:id")
		bookGroup.POST("/add-book")
		bookGroup.DELETE("/delete/:id")
	}

	return &Server{
		Host:   host,
		router: r,
	}

}

func (s *Server) Run() error {
	if err := s.router.Run(s.Host); err != nil {
		return err
	}
	return nil
}
