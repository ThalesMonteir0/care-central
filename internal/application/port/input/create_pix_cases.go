package input

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type CreatePixService interface {
	CreatePix(domain.CreatePixDomain) *rest_err.RestErr
}
