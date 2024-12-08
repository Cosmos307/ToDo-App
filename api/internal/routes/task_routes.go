package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, h *handlers.TaskHandler) {
	r.GET("/tasks/:taskID", h.GetTaskByID)
	r.POST("/tasks", h.CreateTask)
	r.PUT("/tasks/:taskID", h.UpdateTaskByID)
	r.DELETE("/tasks/:taskID", h.DeleteTaskByID)
}
