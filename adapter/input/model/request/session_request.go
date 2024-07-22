package request

import "time"

type SessionRequest struct {
	PatientID     int       `json:"patient_id" biding:"required,min=1"`
	ClinicID      int       `json:"clinic_id" biding:"required,min=1"`
	Paid          bool      `json:"paid" biding:"required"`
	SessionReport string    `json:"session_report"`
	Obs           string    `json:"obs"`
	ValueSession  float64   `json:"value_session" biding:"required,min=0"`
	DtSession     time.Time `json:"dt_session" biding:"required"`
}

type SessionUpdateRequest struct {
	Paid         bool      `json:"paid" biding:"required"`
	ValueSession float64   `json:"value_session" biding:"required,min=0"`
	DtSession    time.Time `json:"dt_session" biding:"required"`
}
