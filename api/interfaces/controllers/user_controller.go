package controllers

import (
	"submane-server/api/interfaces/controllers/database"
	"submane-server/api/usecase"
)

type UserController struct {
	interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				sqlHandler: sqlHandler,
			},
		},
	}
}
