package seeds

import (
	"api/domain"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) (err error) {
	user := domain.User{
		ID: 1,
		Name: "テストユーザ",
		Email: "test@example.com",
	}
	user.SetUid("p7RHrnaOm2fE9QZtbRKlHKQ30Rs1")
	if err = db.Create(&user).Error; err != nil {
		return
	}
	return
}
