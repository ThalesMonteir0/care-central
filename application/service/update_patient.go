package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"strings"
)

func (p *patientService) UpdatePatient(domain domain.PatientDomain) *rest_err.RestErr {
	patient, _ := p.repository.FindPatientByID(domain)
	if patient == nil {
		return rest_err.NewBadRequestError("patient not found")
	}

	patient.Name = strings.ToUpper(domain.Name)
	patient.ResponsibleCPF = domain.ResponsibleCPF
	patient.DateOfBirth = domain.DateOfBirth

	if err := p.repository.UpdatePatient(*patient); err != nil {
		return err
	}

	return nil
}
