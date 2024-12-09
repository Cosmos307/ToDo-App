package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	categoryGroup := r.Group("/categories")

	categoryGroup.GET("/:categoryID", h.GetCategoryByID)
	categoryGroup.POST("/", h.CreateCategory)
	categoryGroup.PUT("/:categoryID", h.UpdateCategoryByID)
	categoryGroup.DELETE("/:categoryID", h.DeleteCategoryByID)
}
