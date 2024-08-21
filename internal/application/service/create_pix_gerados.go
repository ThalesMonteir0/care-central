package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *pixGeradosService) CreatePixGerado(domain domain.PixGeradosDomain) *rest_err.RestErr {
	if err := p.repository.CreatePixGerados(domain); err != nil {
		return err
	}

	return nil
}
