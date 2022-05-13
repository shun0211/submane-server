package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"

	"github.com/labstack/echo/v4"
)

type SubscriptionController struct {
	Interactor usecase.SubscriptionInteractor
}

func NewSubscriptionController(sqlHandler database.SqlHandler) *SubscriptionController {
	return &SubscriptionController{
		Interactor: usecase.SubscriptionInteractor{
			SubscriptionRepository: &database.SubscriptionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SubscriptionController) Index(c echo.Context) (err error) {
	subscriptions, err := controller.Interactor.Subscriptions()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, subscriptions)
	return
}

func(controller *SubscriptionController) Create(c echo.Context) (err error) {
	s := domain.Subscription{}
	c.Bind(&s)
	subscription, err := controller.Interactor.Add(s)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, subscription)
	return
}
