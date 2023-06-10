package main

import (
	"log"

	"github.com/antonpodkur/Blog/config"
	"github.com/antonpodkur/Blog/internal/server"
	"github.com/antonpodkur/Blog/pkg/db"
	"github.com/gin-gonic/gin"
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

    mongoClient := db.MongoClient(cfg)

    router := gin.Default()

    s := server.NewServer(cfg, mongoClient, router)
    s.Run()
}
