package repository

import "github.com/Cosmos307/todo-app/api/internal/models"

type TaskRepository interface {
	GetTasksByUserID(userID int) ([]models.Task, error)
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTaskByID(id int) error
}
