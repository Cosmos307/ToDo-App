package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Cosmos307/todo-app/api/internal/models"
	"github.com/Cosmos307/todo-app/api/internal/repository"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTasksByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var tasks []models.Task
	tasks, err = h.repo.GetTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	task, err := h.repo.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	createdTask, err := h.repo.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if task.ID != 0 && task.ID != taskID {
		log.Println("ID in JSON body does not match URL ID. Ignoring body ID.")
	}
	task.ID = taskID
	err = h.repo.UpdateTaskByID(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = h.repo.DeleteTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
