package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConverterSessionEntityToDomain(session entity.Session) *domain.SessionDomain {
	return &domain.SessionDomain{
		ID:            int(session.ID),
		PatientID:     session.PatientID,
		ClinicID:      session.ClinicID,
		Paid:          session.Paid,
		SessionReport: session.SessionReport,
		Obs:           session.Obs,
		ValueSession:  session.ValueSession,
		DtSession:     *session.DtSession,
	}
}
