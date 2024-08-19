package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *patientRepository) UpdatePatient(domain domain.PatientDomain) *rest_err.RestErr {
	patientEntity := converter.ConvertPatientDomainToEntity(domain)

	if result := p.db.Save(&patientEntity); result.Error != nil {
		return rest_err.NewInternalServerError("unable update patient")
	}

	return nil
}
