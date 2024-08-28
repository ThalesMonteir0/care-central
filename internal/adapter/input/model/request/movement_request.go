package request

type MovementRequest struct {
	ClinicID       int     `json:"clinic_id" biding:"required,min=1"`
	MovementTypeID int     `json:"movement_type_id" biding:"required,min=1"`
	Value          float64 `json:"value" biding:"required"`
	Description    string  `json:"description" biding:"required"`
}

type MovementsFilters struct {
	Type      string `json:"type"`
	DtInitial string `json:"dt_initial"`
	DtFinal   string `json:"dt_final"`
}
