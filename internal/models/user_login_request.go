package models

type UserLoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
