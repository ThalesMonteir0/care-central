package input

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type PixGeradosService interface {
	CreatePixGerado(domain.PixGeradosDomain) *rest_err.RestErr
}
