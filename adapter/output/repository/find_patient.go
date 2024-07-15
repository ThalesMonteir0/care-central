package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (p *patientRepository) FindPatients(domain domain.PatientDomain) ([]domain.PatientDomain, *rest_err.RestErr) {
	var patients []entity.Patient
	patientEntity := converter.ConvertPatientDomainToEntity(domain)

	if result := p.db.Find(&patients, "clinic_id = ?", patientEntity.ClinicID); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get patients")
	}

	return converter.ConvertPatientsEntitiesToDomains(patients), nil
}
