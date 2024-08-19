package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) output.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
