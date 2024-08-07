package entity

import "gorm.io/gorm"

type Movement struct {
	gorm.Model
	ClinicID       int
	Clinic         Clinic
	MovementTypeID int
	MovementsTypes MovementsTypes
	Value          float64
	Description    string
}
