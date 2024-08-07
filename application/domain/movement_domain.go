package domain

import "time"

type MovementDomain struct {
	ID             int
	ClinicID       int
	TypeMovementID int
	Value          float64
	Description    string
	CreatedAt      time.Time
}
