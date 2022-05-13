package domain

import (
	"time"

	"gorm.io/gorm"
)

type Subscriptions []Subscription

type Subscription struct {
	// NOTE: IDフィールドは自動で主キーとして扱われる
	gorm.Model
	Name string
	Price int
	ContractAt time.Time
	UpdateAt time.Time
	UserID int
	User User
}
