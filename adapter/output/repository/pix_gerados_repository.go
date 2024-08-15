package repository

import (
	"github.com/ThalesMonteir0/care-central/application/port/output"
	"gorm.io/gorm"
)

type pixGeradosRepository struct {
	db *gorm.DB
}

func NewPixGeradosRepository(db *gorm.DB) output.PixGeradosRepository {
	return &pixGeradosRepository{
		db: db,
	}
}
