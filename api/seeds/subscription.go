package seeds

import (
	"api/domain"
	"time"

	"gorm.io/gorm"
)

func CreateSubscriptions(db *gorm.DB) (err error) {
	for i := 0; i < 21; i++ {
		subscription := domain.Subscription{
			Name: "テスト",
			Price: 1000,
			ContractedAt: domain.MyTime{Time: time.Now()},
			UserID: 1,
		}
		if err = db.Create(&subscription).Error; err != nil {
			print(err)
			return
		}
	}
	return
}
