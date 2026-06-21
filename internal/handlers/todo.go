package handlers

import (
	"net/http"
	"secure-todo/internal/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Todos *db.TodoRepository
}

type createTodoRequest struct {
	Content string `json:"content"`
	DOW     int    `json:"dow"`
}

type updateTodoRequest struct {
	Content   string `json:"content"`
	DOW       int    `json:"dow"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req createTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, _ := c.Get("userID")
	todoID, err := h.Todos.CreateTodo(userID.(int), req.Content, req.DOW)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": todoID})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("userID")
	todos, err := h.Todos.GetTodosByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Update(c *gin.Context) {
	var req updateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
	}

	userID, _ := c.Get("userID")
	err = h.Todos.UpdateTodo(userID.(int), todoID, req.Content, req.DOW, req.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo successfully updated"})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
	}

	userID, _ := c.Get("userID")
	err = h.Todos.DeleteTodo(userID.(int), todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't delete todo"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo successfully deleted"})
}
