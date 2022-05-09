package domain

import (
	"time"

	"gorm.io/gorm"
)

type Subscriptions []Subscription

type Subscription struct {
	gorm.Model
	// NOTE: IDフィールドは自動で主キーとして扱われる
	ID int
	Name string
	Price int
	ContractDate time.Time
	UpdateDate time.Time
	UserId int
	User User
}
