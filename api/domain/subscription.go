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
	Name string               `json:"name" validate:"required" jaFieldName:"サブスクリプション名"`
	Price int                 `json:"price" validate:"required" jaFieldName:"月額料金"`
	ContractAt MyTime         `json:"contractAt" jaFieldName:"契約日"`
	UserID int                `json:"userId" validate:"required"`
	// HACK:" Key: 'Subscription.User.Email' Error:Field validation for 'Email' failed on the 'required'となるので、一旦バリデーション無視
	User User                 `validate:"-" gorm:"foreignKey:UserID"`
}

type MyTime struct {
	time.Time
}

func (t *MyTime) Scan(value interface{}) error {
	tm, ok := value.(time.Time)
	if !ok {
		// NOTE: Error型を返す
		return fmt.Errorf("error")
	}
	t.Time = tm
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006/01/02"))
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
