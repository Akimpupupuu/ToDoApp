package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Akimpupupuu/ToDoApp/internal/core/config"
	core_logger "github.com/Akimpupupuu/ToDoApp/internal/core/logger"
	core_pgx_pool "github.com/Akimpupupuu/ToDoApp/internal/core/repository/posgres/pool/pgx"
	core_http_middleware "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/middleware"
	core_http_server "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/server"
	statistics_postgres_repository "github.com/Akimpupupuu/ToDoApp/internal/features/statistics/repository/postgres"
	statistics_service "github.com/Akimpupupuu/ToDoApp/internal/features/statistics/service"
	statistics_transport_http "github.com/Akimpupupuu/ToDoApp/internal/features/statistics/transport/http"
	task_postgres_repository "github.com/Akimpupupuu/ToDoApp/internal/features/tasks/repository/postgres"
	tasks_service "github.com/Akimpupupuu/ToDoApp/internal/features/tasks/service"
	tasks_transport_http "github.com/Akimpupupuu/ToDoApp/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/Akimpupupuu/ToDoApp/internal/features/users/repository/postgres"
	users_service "github.com/Akimpupupuu/ToDoApp/internal/features/users/service"
	users_transport_http "github.com/Akimpupupuu/ToDoApp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application timezone", zap.Any("zone", time.Local))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := task_postgres_repository.NewTasksRepository(pool)
	tasksServise := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTaskHTTPHandler(tasksServise)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsServise := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsServise)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(statisticsTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
