package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewSessionsRepository(db *gorm.DB, logger *zap.Logger) output.SessionRepositoryInterface {
	return &sessionRepository{
		db:     db,
		logger: logger,
	}
}
