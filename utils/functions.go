package utils

import "golang-mygram/model/domain"

func UserResponseFunc(user domain.User) domain.UserResponse {
	userResponse := domain.UserResponse{}
	userResponse.ID = user.ID
	userResponse.Age = user.Age
	userResponse.Email = user.Email
	userResponse.Username = user.Username

	return userResponse
}
