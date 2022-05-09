package usecase

import "api/domain"

type SubscriptionInteractor struct {
	SubscriptionRepository SubscriptionRepository
}

func (interactor *SubscriptionInteractor) Subscriptions() (subscriptions domain.Subscriptions, err error) {
	subscriptions, err = interactor.SubscriptionRepository.FindAll()
	return
}
