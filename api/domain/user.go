package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users []User

type User struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Uid []byte `json:"uid" validate:"required"`
}

func (user *User) SetUid(uid string) {
		// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hashedUid, _ := bcrypt.GenerateFromPassword([]byte(uid), 12)
	user.Uid = hashedUid
}

func(user *User) CompareUid(uid string) error {
	return bcrypt.CompareHashAndPassword(user.Uid, []byte(uid))
}
