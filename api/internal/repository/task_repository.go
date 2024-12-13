package repository

import "github.com/Cosmos307/todo-app/api/internal/models"

type TaskRepository interface {
	GetTasksByUserID(userID int) []models.Task
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTaskByID(task *models.Task) (*models.Task, error)
	DeleteTaskByID(id int) error
}
