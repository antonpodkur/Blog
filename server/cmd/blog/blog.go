package main

import (
	"github.com/antonpodkur/Blog/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    s := server.NewServer(router)
    s.Run()
}
