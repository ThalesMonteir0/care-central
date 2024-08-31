package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type patientRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPatientRepository(db *gorm.DB, logger *zap.Logger) output.PatientRepositoryInterface {
	return &patientRepository{
		db:     db,
		logger: logger,
	}
}
