package entity

import "gorm.io/gorm"

type Movement struct {
	gorm.Model
	ClinicID       int
	Clinic         Clinic
	MovementTypeID int
	MovementType   MovementType
	Value          float64
	Description    string
}
