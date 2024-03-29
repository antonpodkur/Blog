package server

import (
	"fmt"
	"log"

	"github.com/antonpodkur/Blog/config"
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	db     *db.Queries
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *db.Queries, router *gin.Engine) *Server {
	return &Server{
		cfg:    cfg,
		db:     db,
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
