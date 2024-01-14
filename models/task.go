package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Priority    string    `json:"priority"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	NotifyTo    string    `json:"notify_to"`
}

func (t *Task) CreateRequestValidate() error {
	if t.Title == "" {
		return fmt.Errorf("title is required")
	}
	if t.Priority == "" {
		return fmt.Errorf("priority is required")
	}
	if t.DueDate.IsZero() {
		return fmt.Errorf("dueDate is required")
	}

	if t.DueDate.Before(time.Now().UTC()) {
		return fmt.Errorf("DueDate must be in the future")
	}

	return nil
}

func (t *Task) UpdateRequestValidate() error {
	if t.DueDate.Before(time.Now().UTC()) {
		return fmt.Errorf("DueDate must be in the future")
	}

	return nil
}