package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) output.SessionRepositoryInterface {
	return &sessionRepository{
		db: db,
	}
}
