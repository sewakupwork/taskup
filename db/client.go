package db

import "taskup/models"

func NewDBClient() (Client, error) {
	// as of now we only support postgres
	return newPostgresClient()
}

type Client interface {
	CreateTask(task *models.Task) error
	DeleteTask(id string) error
	GetTask(id string) (*models.Task, error)
	ListTask() ([]models.Task, error)
	UpdateTask(task *models.Task) error
}
