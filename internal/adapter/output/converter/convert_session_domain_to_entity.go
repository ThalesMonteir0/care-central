package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"gorm.io/gorm"
)

func ConverterSessionDomainToEntity(sessionDomain domain.SessionDomain) *entity.Session {
	return &entity.Session{
		Model:         gorm.Model{ID: uint(sessionDomain.ID)},
		DtSession:     &sessionDomain.DtSession,
		PatientID:     sessionDomain.PatientID,
		ClinicID:      sessionDomain.ClinicID,
		Paid:          sessionDomain.Paid,
		SessionReport: sessionDomain.SessionReport,
		Obs:           sessionDomain.Obs,
		ValueSession:  sessionDomain.ValueSession,
	}
}
