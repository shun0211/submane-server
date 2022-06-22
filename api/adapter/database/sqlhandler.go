package database

import (
	"api/dto"

	"gorm.io/gorm"
)

type SqlHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Exec(string, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Paginate(dto.Page) *gorm.DB
}
