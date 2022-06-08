package domain

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Subscriptions []Subscription

type Subscription struct {
	// NOTE: IDフィールドは自動で主キーとして扱われる
	ID        int             `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time      `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *time.Time      `json:"deletedAt"`
	Name string               `json:"name" validate:"required"`
	Price int                 `json:"price" validate:"required"`
	ContractAt MyTime         `json:"contractAt"`
	UserID int                `json:"userId" validate:"required"`
	// HACK:" Key: 'Subscription.User.Email' Error:Field validation for 'Email' failed on the 'required'となるので、一旦バリデーション無視
	User User                 `validate:"-" gorm:"foreignKey:UserID"`
}

type MyTime struct {
	time.Time
}

func (t *MyTime) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(output), nil
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	layout := "2006-01-02"
	input := strings.Trim(string(data), "\"")
	timeTime, err := time.Parse(layout, input)
	if err != nil {
		return err
	}

	*t = MyTime{timeTime}
	return err
}
