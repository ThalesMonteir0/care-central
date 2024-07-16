package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (p *patientService) DeletePatient(domain domain.PatientDomain) *rest_err.RestErr {
	if err := p.repository.DeletePatient(domain); err != nil {
		return err
	}

	return nil
}
