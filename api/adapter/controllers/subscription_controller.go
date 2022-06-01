package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"
	"strconv"

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

func (controller *SubscriptionController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	subscription, err := controller.Interactor.SubscriptionById(id)
	if err != nil {
		c.JSON(404, NewError(err))
		return
	}

	c.JSON(200, subscription)
	return
}

func(controller *SubscriptionController) Create(c echo.Context) (err error) {
	s := domain.Subscription{}
	c.Bind(&s)
	if err = c.Validate(s); err != nil {
		c.JSON(400, err.Error())
		return
	}
	subscription, err := controller.Interactor.Add(s)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, subscription)
	return
}

func (controller *SubscriptionController) Save(c echo.Context) (err error) {
	s := domain.Subscription{}
	c.Bind(&s)
	if err = c.Validate(s); err != nil {
		c.JSON(400, NewError(err))
		return
	}

	subscription, err := controller.Interactor.Update(s)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(200, subscription)
	return
}

func(controller *SubscriptionController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	subscription := domain.Subscription{
		ID: id,
	}
	err = controller.Interactor.DeleteById(subscription)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, subscription)
	return
}
