package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Akimpupupuu/ToDoApp/internal/core/domain"
	core_logger "github.com/Akimpupupuu/ToDoApp/internal/core/logger"
	core_http_request "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/request"
	core_http_response "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated              int      `json:"tasks_created"`
	TaskCompleted             int      `json:"tasks_crompleted"`
	TaskCompletedRate         *float64 `json:"task_completed_date"`
	TaskAverageCompletionTime *string  `json:"task_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, from, to, err := GetUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/from/to query params")
		return
	}

	statisticsDomains, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}

	response := toDTOFromDomain(statisticsDomains)
	responseHandler.JsonResponse(response, http.StatusOK)
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TaskAverageCompletionTime != nil {
		duration := statistics.TaskAverageCompletionTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:              statistics.TasksCreated,
		TaskCompleted:             statistics.TaskCompleted,
		TaskCompletedRate:         statistics.TaskCompletedRate,
		TaskAverageCompletionTime: avgTime,
	}
}

func GetUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}
