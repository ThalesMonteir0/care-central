package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertPatientDomainToEntity(patientDomain domain.PatientDomain) entity.Patient {
	patient := entity.Patient{
		Name:           patientDomain.Name,
		CpfResponsible: patientDomain.ResponsibleCPF,
		DateOfBirth:    &patientDomain.DateOfBirth,
		ClinicID:       patientDomain.ClinicID,
	}
	patient.ID = uint(patientDomain.ID)

	return patient
}
