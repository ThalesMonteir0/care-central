package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConverterPatientEntityToDomain(patient entity.Patient) *domain.PatientDomain {
	return &domain.PatientDomain{
		ID:             int(patient.ID),
		Name:           patient.Name,
		ResponsibleCPF: patient.CpfResponsible,
		DateOfBirth:    *patient.DateOfBirth,
		ClinicID:       patient.ClinicID,
	}
}
