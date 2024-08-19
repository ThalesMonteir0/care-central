package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"gorm.io/gorm"
)

type movementRepository struct {
	db *gorm.DB
}

func NewMovementRepository(db *gorm.DB) output.MovementRepositoryInterface {
	return &movementRepository{
		db: db,
	}
}
