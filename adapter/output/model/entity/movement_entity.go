package entity

import "gorm.io/gorm"

type Movement struct {
	gorm.DB
	clinicID       int
	Clinic         Clinic
	MovementTypeID int
	MovementsTypes MovementsTypes
	Value          float64
	Description    string
}
