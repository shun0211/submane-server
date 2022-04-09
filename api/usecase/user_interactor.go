package usecase

import (
	"./../domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}


func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	// Note: 変数の初期化は戻り値に指定した変数でされているのでここでは=を使っている
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(domain.User(u))
	return
}

func (interactor *UserInteractor) DeleteById(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.DeleteById(domain.User(u))
	return
}
