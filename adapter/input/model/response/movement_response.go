package response

import "time"

type MovementResponse struct {
	ID             int       `json:"id"`
	ClinicID       int       `json:"clinic_id"`
	MovementTypeID int       `json:"movement_type_id"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
}
