package input

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type PixGeradosService interface {
	CreatePixGerado(domain.PixGeradosDomain) *rest_err.RestErr
}
