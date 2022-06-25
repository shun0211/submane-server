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

type SubscriptionDto struct {
	Page          domain.Page          `json:"page"`
	Subscriptions domain.Subscriptions `json:"subscriptions"`
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

// HACK: すべてでverifyCookieをしているので共通化したい
func (controller *SubscriptionController) Index(c echo.Context) (err error) {
	_, err = verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	userId, _ := strconv.Atoi(c.QueryParam("userId"))
	totalCount, _ := controller.Interactor.GetTotalCount(userId)
	page := CreateToPage(c, totalCount)
	subscriptions, err := controller.Interactor.Subscriptions(userId, page)
	if err != nil {
		c.JSON(500, NewError(err.Error(), ""))
		return
	}
	c.JSON(200, SubscriptionDto{Subscriptions: subscriptions, Page: page})
	return
}

func (controller *SubscriptionController) Show(c echo.Context) (err error) {
	_, err = verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	subscription, err := controller.Interactor.SubscriptionById(id)
	if err != nil {
		c.JSON(404, NewError(err.Error(), ""))
		return
	}

	c.JSON(200, subscription)
	return
}

func (controller *SubscriptionController) Create(c echo.Context) (err error) {
	_, err = verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	s := domain.Subscription{}
	c.Bind(&s)

	if err = c.Validate(s); err != nil {
		messages := GetErrorMessages(err)
		c.JSON(400, messages)
		return
	}

	subscription, err := controller.Interactor.Add(s)
	if err != nil {
		c.JSON(500, NewError(err.Error(), ""))
		return
	}
	c.JSON(201, subscription)
	return
}

func (controller *SubscriptionController) Save(c echo.Context) (err error) {
	_, err = verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	subscription, err := controller.Interactor.SubscriptionById(id)
	if err != nil {
		c.JSON(500, NewError(err.Error(), ""))
	}

	c.Bind(&subscription)
	if err = c.Validate(subscription); err != nil {
		messages := GetErrorMessages(err)
		c.JSON(400, messages)
		return
	}

	subscription, err = controller.Interactor.Update(subscription)
	if err != nil {
		c.JSON(500, NewError(err.Error(), ""))
		return
	}

	c.JSON(200, subscription)
	return
}

func (controller *SubscriptionController) Delete(c echo.Context) (err error) {
	_, err = verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	subscription := domain.Subscription{
		ID: id,
	}
	err = controller.Interactor.DeleteById(subscription)
	if err != nil {
		c.JSON(500, NewError(err.Error(), ""))
		return
	}
	c.JSON(200, subscription)
	return
}
