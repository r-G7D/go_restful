package repository

import (
	"r-G7D/go_restful/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(db *gorm.DB, user domain.User) domain.User
	Update(db *gorm.DB, user domain.User) domain.User
	Delete(db *gorm.DB, user domain.User)
	FindById(db *gorm.DB, userId int) (domain.User, error)
	FindAll(db *gorm.DB) []domain.User
}
