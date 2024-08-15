package http_client

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/application/port/output"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type httpClientPix struct {
}

func NewHttpClientPix() output.HttpClientPix {
	return &httpClientPix{}
}

func (h *httpClientPix) CreatePixCob(domain domain.CreatePixDomain) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (h *httpClientPix) AuthPix() *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}
