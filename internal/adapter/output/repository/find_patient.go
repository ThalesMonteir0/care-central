package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *patientRepository) FindPatients(domain domain.PatientDomain) ([]domain.PatientDomain, *rest_err.RestErr) {
	var patients []entity.Patient
	patientEntity := converter.ConvertPatientDomainToEntity(domain)

	if result := p.db.Find(&patients, "clinic_id = ?", patientEntity.ClinicID); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get patients")
	}

	return converter.ConvertPatientsEntitiesToDomains(patients), nil
}

func (p *patientRepository) FindPatientByID(domain domain.PatientDomain) (*domain.PatientDomain, *rest_err.RestErr) {
	var patient entity.Patient
	patientEntity := converter.ConvertPatientDomainToEntity(domain)

	if result := p.db.First(&patient, patientEntity.ID); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get patient")
	}

	return converter.ConverterPatientEntityToDomain(patient), nil
}
