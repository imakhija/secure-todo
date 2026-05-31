package main

import (
	"fmt"
	"log"
	"net/http"
	"secure-todo/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	router := gin.Default()

	router.StaticFile("/", "./web/index.html")
	router.Static("/static", "./web")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	fmt.Println("Starting server...")
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
