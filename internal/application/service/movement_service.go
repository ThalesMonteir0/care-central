package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
)

type movementService struct {
	repository output.MovementRepositoryInterface
	logger     *zap.Logger
}

func NewMovementService(repository output.MovementRepositoryInterface, logger *zap.Logger) input.MovementServiceInterface {
	return &movementService{
		repository: repository,
		logger:     logger,
	}
}
