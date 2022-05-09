package database

import "api/domain"

type SubscriptionRepository struct {
	SqlHandler
}

func (repo *SubscriptionRepository) FindAll() (subscriptions domain.Subscriptions, err error) {
	if err = repo.Find(&subscriptions).Error; err != nil {
		return
	}
	return
}
