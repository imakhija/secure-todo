package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "./web")

	fmt.Println("Starting server...")
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
