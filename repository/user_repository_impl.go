package repository

import (
	"r-G7D/go_restful/domain"

	"gorm.io/gorm"
)

type UserRepositoryImp struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (r *UserRepositoryImp) Save(db *gorm.DB, user domain.User) domain.User {
	db.Create(&user)
	return user
}

func (r *UserRepositoryImp) Update(db *gorm.DB, user domain.User) domain.User {
	db.Save(&user)
	return user
}

func (r *UserRepositoryImp) Delete(db *gorm.DB, user domain.User) {
	db.Delete(&user)
}

func (r *UserRepositoryImp) FindById(db *gorm.DB, userId int) (domain.User, error) {
	var user domain.User
	result := db.First(&user, userId)
	// result := db.Find(&user, userId)
	return user, result.Error
}

func (r *UserRepositoryImp) FindAll(db *gorm.DB) []domain.User {
	var users []domain.User
	db.Find(&users)
	return users
}
