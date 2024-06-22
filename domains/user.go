package domains

import (
	"github.com/fadhlinw/clean-gin/dto"
	"github.com/fadhlinw/clean-gin/models"
	"github.com/fadhlinw/clean-gin/utils"
	"gorm.io/gorm"
)

type UserService interface {
	WithTrx(trxHandle *gorm.DB) UserService
	GetOneUser(id int) (*models.User, error)
	GetOneUserById(id int) (dto.UserResponseDto, error)
	GetListUsers() ([]dto.UserResponseDto, error)
	GetAllUser(searchQuery string,pagination utils.Pagination) (utils.Pagination, error)
	CreateUser(createUserDto dto.CreateUserRequest) error
	UpdateUser(id uint, dto dto.CreateUserRequest) error
	DeleteUser(id uint) error
	GetOneByEmail(email string) (*models.User, error)
	GetOneUserByEmail(email string) (dto.UserResponseDto, error)
}
