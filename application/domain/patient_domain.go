package domain

import (
	"strings"
	"time"
)

type PatientDomain struct {
	ID             int
	Name           string
	ResponsibleCPF string
	DateOfBirth    time.Time
	ClinicID       int
}

func (p *PatientDomain) ToUpperCaseName() {
	p.Name = strings.ToUpper(p.Name)
}
