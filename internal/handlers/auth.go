package handlers

import (
	"net/http"
	"secure-todo/internal/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Users *db.UserRepository
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't hash password"})
		return
	}

	userID, err := h.Users.CreateUser(req.Username, string(hashedPW))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't register new user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userID})
}
