package controllers

import (
	"api/adapter/database"
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
