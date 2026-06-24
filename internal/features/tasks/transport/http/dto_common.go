package tasks_transport_http

import (
	"time"

	"github.com/Akimpupupuu/ToDoApp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	UserId      int        `json:"user_id"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:          task.ID,
		Version:     task.Version,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,
		UserId:      task.UserId,
	}
}

func tasksDTOFromDomain(tasks []domain.Task) []TaskDTOResponse {
	taskDto := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		taskDto[i] = taskDTOFromDomain(task)
	}

	return taskDto
}
