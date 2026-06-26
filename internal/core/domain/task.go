package domain

import (
	"fmt"
	"time"

	core_errors "github.com/Akimpupupuu/ToDoApp/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	UserId      int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserId int,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		UserId:      authorUserId,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserId int,
) Task {
	return Task{
		ID:          UnitializaedID,
		Version:     UnitializaedVersion,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		UserId:      authorUserId,
	}
}

func (t *Task) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)
	return &duration
}

func (t *Task) Validate() error {
	titleLength := len([]rune(t.Title))
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf("invalid 'Title' length: %d: %w", titleLength, core_errors.ErrInvalidArgument)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("invalid 'Description' length: %d: %w", descriptionLength, core_errors.ErrInvalidArgument)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("'CompletedAt' cant be 'nil' if 'Completed' == 'true': %w", core_errors.ErrInvalidArgument)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("'CompletedAt' cant be before 'CompletedAt': %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("'CompletedAt' must be 'nil' if 'Completed' == 'false': %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(title Nullable[string], description Nullable[string], completed Nullable[bool]) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title' cant be null: %w", core_errors.ErrInvalidArgument)
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' cant be null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = *&patch.Description.Value
	}

	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := t.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*t = tmp

	return nil
}
