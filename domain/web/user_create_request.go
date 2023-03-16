package web

type UserCreateRequest struct {
	ID    int    `json:"id"`
	Name  string `validate:"required,max=100,min=1" json:"name"`
	Email string `validate:"required,email,max=100,min=1" json:"email"`
}
