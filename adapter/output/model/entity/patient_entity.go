package entity

import (
	"gorm.io/gorm"
	"time"
)

type Patient struct {
	gorm.Model
	Name           string
	CpfResponsible string
	DateOfBirth    *time.Time
	ClinicID       int
}
