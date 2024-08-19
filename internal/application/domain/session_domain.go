package domain

import "time"

type SessionDomain struct {
	ID            int
	PatientID     int
	ClinicID      int
	Paid          bool
	SessionReport string
	Obs           string
	ValueSession  float64
	DtSession     time.Time
}
