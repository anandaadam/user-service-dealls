package models

import "time"

type User struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `gorm:"column:username;unique;not null"`
	FullName  string    `gorm:"column:full_name;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	Phone     *string   `gorm:"column:phone;unique"`
	Password  string    `gorm:"column:password"`
	BirthDate time.Time `gorm:"column:birth_date;not null"`
	City      string    `gorm:"column:city;not null"`
	Bio       string    `gorm:"column:bio; not null"`
	IsPremium bool      `gorm:"column:is_premium;default:false"`
	Status    bool      `gorm:"column:status;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "user.users"
}
