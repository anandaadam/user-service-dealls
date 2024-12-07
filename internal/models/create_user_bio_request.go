package models

type CreateUserBioRequest struct {
	FullName  string `validate:"required,fullname"`
	Phone     string
	BirthDate string
	City      string
	Bio       string
}
