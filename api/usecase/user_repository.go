package usecase

import "submane-server/api/domain"

type UserRepository interface {
	FindById(id int) (domain.User, error)
}
