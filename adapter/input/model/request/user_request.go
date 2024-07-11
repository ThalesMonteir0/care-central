package request

type UserRequest struct {
	Name     string `json:"name" biding:"required"`
	Email    string `json:"email" biding:"required"`
	Password string `json:"password" biding:"required,min=6"`
	ClinicID int    `json:"clinic_id" biding:"required"`
}
