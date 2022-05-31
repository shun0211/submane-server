package domain

import (
	"time"
)

type Subscriptions []Subscription

type Subscription struct {
	// NOTE: IDフィールドは自動で主キーとして扱われる
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	Name string           `json:"name" validate:"required"`
	Price int             `json:"price" validate:"required"`
	ContractAt time.Time  `json:"contract_at" validate:"required"`
	UserID int            `json:"userId" validate:"required"`
	// HACK:" Key: 'Subscription.User.Email' Error:Field validation for 'Email' failed on the 'required'となるので、一旦バリデーション無視
	User User             `validate:"-"`
}
