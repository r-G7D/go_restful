package web

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
}
