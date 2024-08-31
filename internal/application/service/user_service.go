package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
)

type userService struct {
	repository output.UserRepositoryInterface
	logger     *zap.Logger
}

func NewUserService(repository output.UserRepositoryInterface, logger *zap.Logger) input.UserServiceInterface {
	return &userService{
		repository: repository,
		logger:     logger,
	}
}
