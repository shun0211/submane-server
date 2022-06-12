package usecase

import "api/domain"

type SubscriptionRepository interface {
	FindById(id int) (domain.Subscription, error)
	FindAll(userId int) (domain.Subscriptions, error)
	Store(domain.Subscription) (domain.Subscription, error)
	Update(domain.Subscription) (domain.Subscription, error)
	DeleteById(domain.Subscription) (error)
}
