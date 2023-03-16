package helper

import (
	"r-G7D/go_restful/domain"
	"r-G7D/go_restful/domain/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
