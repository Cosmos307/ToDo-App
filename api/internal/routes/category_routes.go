package routes

import (
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	r.GET("/categorys/:categoryID", h.GetCategoryByID)
	r.POST("/categorys", h.CreateCategory)
	r.PUT("/categorys/:categoryID", h.UpdateCategoryByID)
	r.DELETE("/categorys/:categoryID", h.DeleteCategoryByID)
}
