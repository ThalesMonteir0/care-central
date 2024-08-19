package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionService) UpdateSession(sessionDomain domain.SessionDomain) *rest_err.RestErr {
	session, _ := s.repository.FindSessionByID(sessionDomain)
	if session == nil || session.ClinicID != sessionDomain.ClinicID {
		return rest_err.NewBadRequestError("session not exists")
	}

	sesssionDomainUpdate := domain.SessionDomain{
		ID:            session.ID,
		PatientID:     session.PatientID,
		ClinicID:      session.ClinicID,
		Paid:          sessionDomain.Paid,
		SessionReport: session.SessionReport,
		Obs:           session.Obs,
		ValueSession:  sessionDomain.ValueSession,
		DtSession:     sessionDomain.DtSession,
	}

	if err := s.repository.UpdateSession(sesssionDomainUpdate); err != nil {
		return err
	}

	return nil
}
