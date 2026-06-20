package main

import (
	"fmt"
	"log"
	"net/http"
	"secure-todo/internal/db"
	"secure-todo/internal/handlers"
	"secure-todo/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := &db.UserRepository{
		Pool: pool,
	}

	authHandler := &handlers.AuthHandler{
		Users: userRepo,
	}

	router := gin.Default()

	router.StaticFile("/", "./web/index.html")
	router.Static("/static", "./web")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	router.POST("/api/register", authHandler.Register)
	router.POST("/api/login", authHandler.Login)

	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"user_id": userID})
	})

	fmt.Println("Starting server...")
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
