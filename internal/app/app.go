package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/AndroDeMohawk/sso-app/internal/app/grpc"
	"github.com/AndroDeMohawk/sso-app/internal/services/auth"
	"github.com/AndroDeMohawk/sso-app/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage := postgres.New(storagePath)

	authService := auth.New(log, storage, storage, storage, tokenTTL)
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
