package service

import (
	"context"
	"r-G7D/go_restful/domain"
	"r-G7D/go_restful/domain/web"
	"r-G7D/go_restful/helper"
	"r-G7D/go_restful/helper/exception"
	"r-G7D/go_restful/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		validate:       validate,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	// Validate the request
	err := s.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	// Convert request to domain
	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	// Save to database
	s.UserRepository.Save(s.DB, user)

	return helper.ToUserResponse(user)
}

func (s *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := s.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	user, err := s.UserRepository.FindById(s.DB, int(request.ID))
	if err != nil {
		panic(exception.NewNotFound("user not found"))
	}

	user.Name = request.Name
	user.Email = request.Email
	//set new data after update

	response := s.UserRepository.Update(s.DB, user)

	return helper.ToUserResponse(response)
}

func (s *UserServiceImpl) Delete(ctx context.Context, userId int) {
	//shorter
	//!somehow when id is being validated, its always null
	// if err := s.validate.Struct(userId); err != nil {
	// 	panic(err)
	// }

	user, err := s.UserRepository.FindById(s.DB, userId)
	if err != nil {
		panic(exception.NewNotFound("user not found"))
	}

	s.UserRepository.Delete(s.DB, user)
}

func (s *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	// Validate the request
	//!somehow when id is being validated, its always null
	//!morover, it works on update but doesnt work on delete and GET by id
	//*apparently, delete and find by id doesnt need validation
	// if err := s.validate.Struct(userId); err != nil {
	// 	panic(err)
	// }

	user, err := s.UserRepository.FindById(s.DB, userId)
	if err != nil {
		panic(exception.NewNotFound("user not found"))
	}

	return helper.ToUserResponse(user)
}

func (s *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	users := s.UserRepository.FindAll(s.DB)

	return helper.ToUserResponses(users)
}
