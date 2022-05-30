package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Users []User

type User struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	Name string           `json:"name"`
	Email string          `json:"email" validate:"required"`
	Uid []byte            `json:"uid" validate:"required"`
}

func (user *User) SetUid(uid string) {
		// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hashedUid, _ := bcrypt.GenerateFromPassword([]byte(uid), 12)
	user.Uid = hashedUid
}

func(user *User) CompareUid(uid string) error {
	return bcrypt.CompareHashAndPassword(user.Uid, []byte(uid))
}
