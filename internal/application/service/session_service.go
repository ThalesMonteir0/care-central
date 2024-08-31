package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
)

type sessionService struct {
	repository output.SessionRepositoryInterface
	logger     *zap.Logger
}

func NewSessionService(repository output.SessionRepositoryInterface, logger *zap.Logger) input.SessionServiceInterface {
	return &sessionService{
		repository: repository,
		logger:     logger,
	}
}
