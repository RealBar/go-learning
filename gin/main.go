package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	V1Api := engine.Group("/api/mediaspace/v1")

	imageApi := V1Api.Group("/image")
	imageApi.Handle()
	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("exit with errpr:" + err.Error())
	}
}
