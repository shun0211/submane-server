package database

import (
	"api/domain"

)

// NOTE: ここでフィールド名を省略しているため、SqlHandlerを省略してFindフィールドなどにアクセスできる
type SubscriptionRepository struct {
	SqlHandler
}

func (repo *SubscriptionRepository) FindById(id int) (subscription domain.Subscription, err error) {
	if err = repo.First(&subscription, id).Error; err != nil {
		return
	}
	return
}

func (repo *SubscriptionRepository) FindAll(userId int, page domain.Page) (subscriptions domain.Subscriptions, err error) {
	if err = repo.Paginate(page).Find(&subscriptions, "user_id = ?", userId).Error; err != nil {
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

func(repo *SubscriptionRepository) Update(s domain.Subscription) (subscription domain.Subscription, err error) {
	if err = repo.Save(&s).Error; err != nil {
		return
	}
	subscription = s
	return
}

func(repo *SubscriptionRepository) DeleteById(s domain.Subscription) (err error) {
	if err = repo.Delete(&s).Error; err != nil {
		return
	}
	return
}

func(repo *SubscriptionRepository) GetTotalCount(userId int) (totalCount int, err error) {
	subscriptions := repo.Find(&domain.Subscriptions{}, "user_id = ?", userId)
	if err = subscriptions.Error; err != nil {
		return
	}
	totalCount = int(subscriptions.RowsAffected)
	return
}
