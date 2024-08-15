package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type HttpClientPix interface {
	CreatePixCob(domain domain.CreatePixDomain) *rest_err.RestErr
	AuthPix() *rest_err.RestErr
}
