package dto

import "api/domain"

type SubscriptionDto struct {
	Page          Page                 `json:"page"`
	Subscriptions domain.Subscriptions `json:"subscriptions"`
}
