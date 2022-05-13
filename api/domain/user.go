package domain

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	Name string `validate:"required"`
	Email string `validate:"required"`
}
