package usecase

import "api/domain"

type SubscriptionInteractor struct {
	SubscriptionRepository SubscriptionRepository
}

func (interactor *SubscriptionInteractor) Subscriptions() (subscriptions domain.Subscriptions, err error) {
	subscriptions, err = interactor.SubscriptionRepository.FindAll()
	return
}

func(interactor *SubscriptionInteractor) Add(s domain.Subscription) (subscription domain.Subscription, err error) {
	subscription, err = interactor.SubscriptionRepository.Store(s)
	return
}
