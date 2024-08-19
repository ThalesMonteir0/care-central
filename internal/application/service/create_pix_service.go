package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type createPixService struct {
	httpClientPix     output.HttpClientPix
	pixGeradosService input.PixGeradosService
}

func NewCreatePixService(httpClientPix output.HttpClientPix, pixGeradosService input.PixGeradosService) input.CreatePixService {
	return &createPixService{
		httpClientPix:     httpClientPix,
		pixGeradosService: pixGeradosService,
	}
}

func (c *createPixService) CreatePix(domain domain.CreatePixDomain) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}
