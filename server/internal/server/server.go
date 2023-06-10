package server

import (
	"fmt"
	"log"

	"github.com/antonpodkur/Blog/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
    cfg *config.Config
    mongoClient *mongo.Client
    router *gin.Engine
}

func NewServer(cfg *config.Config, mongoClient *mongo.Client, router *gin.Engine) *Server {
    return &Server{
        cfg: cfg,
        mongoClient: mongoClient,
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
