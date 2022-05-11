package usecase

import "api/domain"

type SubscriptionRepository interface {
	FindAll() (domain.Subscriptions, error)
	Store(domain.Subscription) (domain.Subscription, error)
}
