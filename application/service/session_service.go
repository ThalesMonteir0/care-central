package service

import (
	"github.com/ThalesMonteir0/care-central/application/port/input"
	"github.com/ThalesMonteir0/care-central/application/port/output"
)

type sessionService struct {
	repository output.SessionRepositoryInterface
}

func NewSessionService(repository output.SessionRepositoryInterface) input.SessionServiceInterface {
	return &sessionService{
		repository: repository,
	}
}
