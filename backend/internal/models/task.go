package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "Pending"
	StatusInProgress TaskStatus = "In-Progress"
	StatusCompleted  TaskStatus = "Completed"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Status      TaskStatus     `json:"status" gorm:"type:varchar(20);not null;default:'Pending'"`
	DueDate     *time.Time     `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Validate performs validation on the task
func (t *Task) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	if t.Status != StatusPending && t.Status != StatusInProgress && t.Status != StatusCompleted {
		return ErrInvalidStatus
	}
	return nil
}

// Custom errors
var (
	ErrEmptyTitle    = errors.New("title cannot be empty")
	ErrInvalidStatus = errors.New("invalid task status")
)
