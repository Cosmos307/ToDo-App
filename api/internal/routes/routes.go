package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	categoryHandler *handlers.CategoryHandler) {

	RegisterCategoryRoutes(r, categoryHandler)
}
