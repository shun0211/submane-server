package usecase

import (
	"api/domain"
)

type SubscriptionInteractor struct {
	SubscriptionRepository SubscriptionRepository
}

func (interactor *SubscriptionInteractor) SubscriptionById(id int) (subscription domain.Subscription, err error) {
	subscription, err = interactor.SubscriptionRepository.FindById(id)
	return
}

func (interactor *SubscriptionInteractor) Subscriptions(userId int, page domain.Page) (subscriptions domain.Subscriptions, err error) {
	subscriptions, err = interactor.SubscriptionRepository.FindAll(userId, page)
	return
}

func(interactor *SubscriptionInteractor) Add(s domain.Subscription) (subscription domain.Subscription, err error) {
	subscription, err = interactor.SubscriptionRepository.Store(s)
	return
}

func(interactor *SubscriptionInteractor) Update(s domain.Subscription) (subscription domain.Subscription, err error) {
	subscription, err = interactor.SubscriptionRepository.Update(s)
	return
}

func(interactor *SubscriptionInteractor) DeleteById(s domain.Subscription) (err error) {
	err = interactor.SubscriptionRepository.DeleteById(s)
	return
}

func(interactor *SubscriptionInteractor) GetTotalCount(userId int) (totalCount int, err error) {
	totalCount, err = interactor.SubscriptionRepository.GetTotalCount(userId)
	return
}
