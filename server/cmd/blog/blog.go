package main

import (
	"database/sql"
	"log"

	"github.com/antonpodkur/Blog/config"
	dbConn "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting api server...")

	configPath := config.GetConfigPath("local")
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("LoadConfig: %v", err.Error())
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("ParseConfig: %v", err.Error())
	}

	conn, err := sql.Open(cfg.Postgres.Driver, cfg.Postgres.Source)
	if err != nil {
		log.Fatal("could not connect to postgres database")
	}

	db := dbConn.New(conn)

	router := gin.Default()

	s := server.NewServer(cfg, db, router)
	s.Run()
}
