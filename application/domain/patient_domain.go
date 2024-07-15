package domain

import "time"

type PatientDomain struct {
	ID             int
	Name           string
	ResponsibleCPF string
	DateOfBirth    time.Time
	ClinicID       int
}
