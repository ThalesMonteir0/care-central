package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *pixGeradosRepository) CreatePixGerados(domain domain.PixGeradosDomain) *rest_err.RestErr {
	pixGeradoEntity := converter.ConvertPixGeradosDomainToEntity(domain)

	if tx := p.db.Create(&pixGeradoEntity); tx.Error != nil {
		return rest_err.NewInternalServerError("unable insert in pixGerados")
	}

	return nil
}
