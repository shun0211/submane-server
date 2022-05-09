package usecase

import "api/domain"

type SubscriptionRepository interface {
	FindAll() (domain.Subscriptions, error)
}
