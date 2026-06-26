package domain

import "time"

type Statistics struct {
	TasksCreated              int
	TaskCompleted             int
	TaskCompletedRate         *float64
	TaskAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	taskCompleted int,
	taskCompletedRate *float64,
	taskAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:              tasksCreated,
		TaskCompleted:             taskCompleted,
		TaskCompletedRate:         taskCompletedRate,
		TaskAverageCompletionTime: taskAverageCompletionTime,
	}
}
