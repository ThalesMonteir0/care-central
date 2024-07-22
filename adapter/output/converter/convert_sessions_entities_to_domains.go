package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConverterSessionsEntitiesToDomains(sessions []entity.Session) []domain.SessionDomain {
	var sessionsDomains []domain.SessionDomain

	for _, session := range sessions {
		sessionDomain := domain.SessionDomain{
			ID:            int(session.ID),
			PatientID:     session.PatientID,
			ClinicID:      session.ClinicID,
			Paid:          session.Paid,
			SessionReport: session.SessionReport,
			Obs:           session.Obs,
			ValueSession:  session.ValueSession,
			DtSession:     *session.DtSession,
		}

		sessionsDomains = append(sessionsDomains, sessionDomain)
	}

	return sessionsDomains
}
