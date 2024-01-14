package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"taskup/models"
)

func newPostgresClient() (Client, error) {
	dbURL := "postgres://testuser:testpass@postgres:5432/taskdb?sslmode=disable"

	ds, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	ds.AutoMigrate(&models.Task{})

	return &psclient{ds: ds}, nil
}

type psclient struct {
	ds *gorm.DB
}

func (c psclient) CreateTask(task *models.Task) error {
	txn := c.ds.Create(task)
	return txn.Error
}

func (c psclient) DeleteTask(id string) error {
	var tsk models.Task
	if err := c.ds.First(&tsk, id).Error; err != nil {
		return err
	}

	return c.ds.Delete(&tsk).Error
}

func (c psclient) GetTask(id string) (*models.Task, error) {
	var tsk models.Task
	if err := c.ds.First(&tsk, id).Error; err != nil {
		return nil, err
	}

	return &tsk, nil

}

func (c psclient) ListTask() ([]models.Task, error) {
	var tasks []models.Task

	if err := c.ds.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c psclient) UpdateTask(task *models.Task) error {
	return c.ds.Save(&task).Error
}
