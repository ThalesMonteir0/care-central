package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type PixGeradosRepository interface {
	CreatePixGerados(domain.PixGeradosDomain) *rest_err.RestErr
}
