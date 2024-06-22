package dto

import "time"

type UserRequestDto struct {
}

type CreateUserRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Age          uint8  `json:"age" binding:"required"`
	Birthday     string `json:"birthday" binding:"required"`
	MemberNumber int    `json:"member_number" binding:"required"`
}

type UserResponseDto struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Age          uint8     `json:"age"`
	Birthday     string    `json:"birthday"`
	MemberNumber int       `json:"member_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
