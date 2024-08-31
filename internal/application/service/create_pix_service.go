package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"go.uber.org/zap"
	"strconv"
)

type createPixService struct {
	httpClientPix     output.HttpClientPix
	pixGeradosService input.PixGeradosService
	logger            *zap.Logger
}

func NewCreatePixService(httpClientPix output.HttpClientPix, pixGeradosService input.PixGeradosService, logger *zap.Logger) input.CreatePixService {
	return &createPixService{
		httpClientPix:     httpClientPix,
		pixGeradosService: pixGeradosService,
		logger:            logger,
	}
}

func (c *createPixService) CreatePix(PixDomain domain.CreatePixDomain) (string, *rest_err.RestErr) {
	token, err := c.httpClientPix.AuthPix()
	if err != nil {
		return "", err
	}

	createPixResponse, err := c.httpClientPix.CreatePixCob(PixDomain, token)
	if err != nil {
		return "", err
	}

	valuePix, errConvert := strconv.ParseFloat(createPixResponse.Valor.Original, 64)
	if errConvert != nil {

	}

	pixGeradosDomain := domain.PixGeradosDomain{
		TxID:      createPixResponse.Txid,
		Value:     valuePix,
		SessionID: PixDomain.SessionID,
	}

	if err := c.pixGeradosService.CreatePixGerado(pixGeradosDomain); err != nil {
		return "", err
	}

	return createPixResponse.PixCopiaECola, nil
}
