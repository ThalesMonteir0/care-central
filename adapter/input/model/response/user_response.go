package response

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ClinicID int    `json:"clinic_id"`
}
