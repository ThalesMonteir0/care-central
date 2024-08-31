package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type movementRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMovementRepository(db *gorm.DB, logger *zap.Logger) output.MovementRepositoryInterface {
	return &movementRepository{
		db:     db,
		logger: logger,
	}
}
