package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) output.UserRepositoryInterface {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
