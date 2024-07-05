package entity

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	SocialReason string
	CpfCnpj      string
	Fantasy      string
}
