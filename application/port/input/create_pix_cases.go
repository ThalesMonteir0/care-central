package input

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type CreatePixService interface {
	CreatePix(domain.CreatePixDomain) *rest_err.RestErr
}
