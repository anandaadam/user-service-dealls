package models

import "time"

type CreateUserBioResponse struct {
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Born      string    `json:"born"`
	Bio       string    `json:"bio"`
}
