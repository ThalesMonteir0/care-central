package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConvertPatientsEntitiesToDomains(patients []entity.Patient) []domain.PatientDomain {
	var patientsDomain []domain.PatientDomain

	for _, patient := range patients {
		patientDomain := domain.PatientDomain{
			ID:             int(patient.ID),
			Name:           patient.Name,
			ResponsibleCPF: patient.CpfResponsible,
			DateOfBirth:    *patient.DateOfBirth,
			ClinicID:       patient.ClinicID,
		}

		patientsDomain = append(patientsDomain, patientDomain)
	}
	return patientsDomain
}
