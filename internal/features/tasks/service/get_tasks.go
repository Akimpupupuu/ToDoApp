package tasks_service

import (
	"context"
	"fmt"

	"github.com/Akimpupupuu/ToDoApp/internal/core/domain"
	core_errors "github.com/Akimpupupuu/ToDoApp/internal/core/errors"
)

func (s *TasksService) GetTasks(ctx context.Context, userID *int, limit *int, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		fmt.Errorf("limit must be non-negativ: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		fmt.Errorf("offset must be non-negativ: %w", core_errors.ErrInvalidArgument)
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return tasks, nil
}
