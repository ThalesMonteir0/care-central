package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type pixGeradosRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPixGeradosRepository(db *gorm.DB, logger *zap.Logger) output.PixGeradosRepository {
	return &pixGeradosRepository{
		db:     db,
		logger: logger,
	}
}
