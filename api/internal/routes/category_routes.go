package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	categoryGroup := r.Group("/categories")

	categoryGroup.GET("/:id", h.GetCategoryByID)
	categoryGroup.POST("/", h.CreateCategory)
	categoryGroup.PUT("/:id", h.UpdateCategory)
	categoryGroup.DELETE("/:id", h.DeleteCategoryByID)
}
