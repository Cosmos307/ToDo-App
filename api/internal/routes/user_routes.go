package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *handlers.UserHandler) {
	r.GET("/users/:userID", h.GetUserByID)
	r.POST("/users", h.CreateUser)
	r.PUT("/users/:userID", h.UpdateUser)
	r.DELETE("/users/:userID", h.DeleteUserByID)
}
