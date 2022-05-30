package database

import "api/domain"

// NOTE: ここでフィールド名を省略しているため、SqlHandlerを省略してFindフィールドなどにアクセスできる
type SubscriptionRepository struct {
	SqlHandler
}

func (repo *SubscriptionRepository) FindAll() (subscriptions domain.Subscriptions, err error) {
	if err = repo.Find(&subscriptions).Error; err != nil {
		return
	}
	return
}

func(repo *SubscriptionRepository) Store(s domain.Subscription) (subscription domain.Subscription, err error) {
	if err = repo.Create(&s).Error; err != nil {
		return
	}
	subscription = s
	return
}
