package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
)

type movementService struct {
	repository output.MovementRepositoryInterface
}

func NewMovementService(repository output.MovementRepositoryInterface) input.MovementServiceInterface {
	return &movementService{
		repository: repository,
	}
}
