package output

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/response"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type HttpClientPix interface {
	CreatePixCob(domain domain.CreatePixDomain, token string) (response.CreatePixEFIResponse, *rest_err.RestErr)
	AuthPix() (string, *rest_err.RestErr)
}
