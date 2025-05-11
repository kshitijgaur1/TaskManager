package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishwa-ai/task-manager/internal/models"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

// GetTasks returns all tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []models.Task
	result := h.db.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask creates a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	// Read the raw request body for logging
	body, err := c.GetRawData()
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Printf("Received request body: %s", string(body))

	// Create a map to parse the JSON
	var taskData map[string]interface{}
	if err := json.Unmarshal(body, &taskData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Create a new task and populate it
	task := models.Task{
		Title:       taskData["title"].(string),
		Description: taskData["description"].(string),
		Status:      models.TaskStatus(taskData["status"].(string)),
	}

	// Handle due_date if present
	if dueDate, ok := taskData["due_date"].(string); ok && dueDate != "" {
		parsedDate, err := time.Parse(time.RFC3339, dueDate)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		task.DueDate = &parsedDate
	}

	if err := task.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Create(&task)
	if result.Error != nil {
		log.Printf("Error creating task: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := h.db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Read the raw request body for logging
	body, err := c.GetRawData()
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Printf("Received request body: %s", string(body))

	// Create a map to parse the JSON
	var taskData map[string]interface{}
	if err := json.Unmarshal(body, &taskData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Update task fields
	task.Title = taskData["title"].(string)
	task.Description = taskData["description"].(string)
	task.Status = models.TaskStatus(taskData["status"].(string))

	// Handle due_date if present
	if dueDate, ok := taskData["due_date"].(string); ok && dueDate != "" {
		parsedDate, err := time.Parse(time.RFC3339, dueDate)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		task.DueDate = &parsedDate
	} else {
		task.DueDate = nil
	}

	if err := task.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Save(&task)
	if result.Error != nil {
		log.Printf("Error updating task: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	result := h.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
