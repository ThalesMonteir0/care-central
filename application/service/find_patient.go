package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (p *patientService) FindPatient(domain domain.PatientDomain) ([]domain.PatientDomain, *rest_err.RestErr) {

	patients, err := p.repository.FindPatients(domain)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
