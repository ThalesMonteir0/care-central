package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
)

type pixGeradosService struct {
	repository output.PixGeradosRepository
	logger     *zap.Logger
}

func NewPixGeradosService(repository output.PixGeradosRepository, logger *zap.Logger) input.PixGeradosService {
	return &pixGeradosService{
		repository: repository,
		logger:     logger,
	}
}
