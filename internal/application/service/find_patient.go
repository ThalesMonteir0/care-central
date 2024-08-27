package service

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (p *patientService) FindPatient(domain domain.PatientDomain, filters request.PatientFilters) ([]domain.PatientDomain, *rest_err.RestErr) {

	patients, err := p.repository.FindPatients(domain, filters)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
