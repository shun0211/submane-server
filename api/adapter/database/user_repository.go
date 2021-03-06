package database

import "api/domain"

// NOTE: ここでフィールド名を省略しているため、SqlHandlerを省略してFindフィールドなどにアクセスできる
type UserRepository struct {
	SqlHandler
}

// NOTE: usecaseでFindByIdなどのインターフェースが定義されているので、外側にあるcontroller(presenter)がインターフェースを実装する
// NOTE: returnではuserとerrが返る
func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = repo.First(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = repo.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user domain.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByEmail(email string) (user domain.User, err error) {
	if err = repo.First(&user, "email = ?", email).Error; err != nil {
		return
	}
	return
}
