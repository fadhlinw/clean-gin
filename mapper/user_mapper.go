package mapper

import (
	"github.com/fadhlinw/clean-gin/dto"
	"github.com/fadhlinw/clean-gin/models"
	"github.com/fadhlinw/clean-gin/utils"
)

func ToUsersResponseDto(users []models.User) []dto.UserResponseDto {
	user := make([]dto.UserResponseDto, len(users))
	for i, item := range users {
		user[i] = ToUserResponseDto(item)
	}
	return user
}

func ToUserResponseDto(user models.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		Age:              user.Age,
		Birthday:         user.Birthday,
		MemberNumber:     user.MemberNumber,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}
}

func ToUserModel(user dto.CreateUserRequest) (models.User, error) {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Name:               user.Name,
		Email:              user.Email,
		Password:           password,
		Age:                user.Age,
		Birthday:           user.Birthday,
		MemberNumber:       user.MemberNumber,
	}, nil
}