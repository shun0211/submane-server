package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"
	"api/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
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
	cookie, err := c.Cookie("userId")
	if err != nil {
		c.JSON(401, NewError(err))
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, NewError(err))
		return
	}

	payload := token.Claims.(*jwt.StandardClaims)
	userId, _ := strconv.Atoi(payload.Subject)
	print(userId)

	s := domain.Subscription{}
	c.Bind(&s)

	if s.UserID != userId {
		c.JSON(400, "Invalid Cookie or userID")
		return
	}

	if err = c.Validate(s); err != nil {
		messages := utils.GetErrorMessages(err)
		c.JSON(400, messages)
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
	id, _ := strconv.Atoi(c.Param("id"))

	subscription, err := controller.Interactor.SubscriptionById(id)
	if err != nil {
		c.JSON(500, NewError(err))
	}

	c.Bind(&subscription)
	if err = c.Validate(subscription); err != nil {
		c.JSON(400, NewError(err))
		return
	}

	subscription, err = controller.Interactor.Update(subscription)
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
