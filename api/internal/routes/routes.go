package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler, categoryHandler *handlers.CategoryHandler) {
	RegisterUserRoutes(r, userHandler)
	RegisterTaskRoutes(r, taskHandler)
	RegisterCategoryRoutes(r, categoryHandler)
}
