package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Role     string
	Active   bool
	Email    string
	ClinicID int
	Clinic   Clinic
}
