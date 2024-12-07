package models

type CreateUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
