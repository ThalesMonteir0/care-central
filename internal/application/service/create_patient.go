package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *patientService) CreatePatient(domain domain.PatientDomain) *rest_err.RestErr {
	domain.ToUpperCaseName()

	if err := p.repository.CreatePatient(domain); err != nil {
		return err
	}

	return nil
}
