package output

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type PixGeradosRepository interface {
	CreatePixGerados(domain.PixGeradosDomain) *rest_err.RestErr
}
