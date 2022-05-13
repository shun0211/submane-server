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
	User User
}
