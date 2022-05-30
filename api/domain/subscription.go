package domain

import (
	"time"

	"gorm.io/gorm"
)

type Subscriptions []Subscription

type Subscription struct {
	// NOTE: IDフィールドは自動で主キーとして扱われる
	gorm.Model
	Name string `validate:"required"`
	Price int `validate:"required"`
	ContractAt time.Time
	UserID int `validate:"required"`
	// HACK:" Key: 'Subscription.User.Email' Error:Field validation for 'Email' failed on the 'required'となるので、一旦バリデーション無視
	User User `validate:"-"`
}
