package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConvertPixGeradosDomainToEntity(pixGeradoDomain domain.PixGeradosDomain) entity.PixGerado {
	return entity.PixGerado{
		ID:        pixGeradoDomain.ID,
		Value:     pixGeradoDomain.Value,
		TxID:      pixGeradoDomain.TxID,
		SessionID: pixGeradoDomain.SessionID,
	}
}
