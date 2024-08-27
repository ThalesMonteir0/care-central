package repository

import (
	"fmt"
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *patientRepository) FindPatients(domain domain.PatientDomain, filters request.PatientFilters) ([]domain.PatientDomain, *rest_err.RestErr) {
	var patients []entity.Patient
	patientEntity := converter.ConvertPatientDomainToEntity(domain)

	where := whereFindPatients(patientEntity.ClinicID, filters)

	if result := p.db.Find(&patients, where); result.Error != nil {
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

func whereFindPatients(clinicID int, filters request.PatientFilters) string {
	where := fmt.Sprintf("clinic_id = %d", clinicID)

	if filters.PatientID != "" {
		where += fmt.Sprintf(" AND id = %s", filters.PatientID)
	}

	return where
}
