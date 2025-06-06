package app

import (
	"context"
	"errors"
	"go1/internal/config"
	delivery "go1/internal/delivery/http"
	"go1/internal/repository"
	"go1/internal/server"
	"go1/internal/service"
	"go1/pkg/database/postgres"
	"go1/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Config: %+v", cfg)

	db, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		logger.Error(err)

		return
	}
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := delivery.NewHandler(services)

	srv := server.NewServer(cfg.Http, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
