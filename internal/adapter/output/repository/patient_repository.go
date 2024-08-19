package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) output.PatientRepositoryInterface {
	return &patientRepository{
		db: db,
	}
}
