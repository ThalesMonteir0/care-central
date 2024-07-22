package entity

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	DtSession     *time.Time
	PatientID     int
	ClinicID      int
	Paid          bool
	SessionReport string
	Obs           string
	ValueSession  float64
}
