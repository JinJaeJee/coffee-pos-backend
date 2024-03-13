package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	router := gin.New()
	router.Use(gin.Logger())
}

