package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
    router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
    return &Server{
        router: router,
    }
}

func (s *Server) Run() {
    err := s.MapHandlers(s.router)
    if err != nil {
        log.Panic(err)
    }

    if err := s.router.Run(":4000"); err != nil {
        log.Panic(err)
    }

    fmt.Println("Server started on port :4000")
}
