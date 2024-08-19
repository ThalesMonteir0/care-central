package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
)

type pixGeradosService struct {
	repository output.PixGeradosRepository
}

func NewPixGeradosService(repository output.PixGeradosRepository) input.PixGeradosService {
	return &pixGeradosService{
		repository: repository,
	}
}
