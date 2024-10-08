package request

import "time"

type PatientRequest struct {
	Name           string    `json:"name" biding:"required"`
	ResponsibleCPF string    `json:"responsible_cpf" biding:"required"`
	ClinicID       int       `json:"clinic_id" biding:"required,min=1"`
	DateOfBirth    time.Time `json:"date_of_birth" biding:"required"`
}

type PatientRequestUpdate struct {
	Name           string    `json:"name"`
	ResponsibleCPF string    `json:"responsible_cpf"`
	DateOfBirth    time.Time `json:"date_of_birth"`
}

type PatientFilters struct {
	PatientID string `json:"patient_id"`
}
