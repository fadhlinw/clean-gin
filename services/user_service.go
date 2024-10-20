package services

import (
	"net/http"

	"github.com/fadhlinw/clean-gin/constants"
	"github.com/fadhlinw/clean-gin/domains"
	"github.com/fadhlinw/clean-gin/dto"
	httperror "github.com/fadhlinw/clean-gin/error"
	"github.com/fadhlinw/clean-gin/lib"
	"github.com/fadhlinw/clean-gin/mapper"
	"github.com/fadhlinw/clean-gin/models"
	"github.com/fadhlinw/clean-gin/repository"
	"github.com/fadhlinw/clean-gin/utils"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
	smtpClient lib.SMTP
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, smtpClient lib.SMTP, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
		smtpClient: smtpClient,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id int) (user *models.User, err error) {
	err = s.repository.First(&user, id).Error
	if err != nil {
		return nil, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusInternalServerError)
	}
	return user, err
}

// GetOneUserById implements domains.UserService.
func (s UserService) GetOneUserById(id int) (dto dto.UserResponseDto, err error) {
	user := models.User{}
	err = s.repository.First(&user, id).Error
	if err != nil {
		return dto, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusInternalServerError)
	}
	return mapper.ToUserResponseDto(user), err
}

// GetListUsers implements domains.UserService.
func (s UserService) GetListUsers() ([]dto.UserResponseDto, error) {
	users := []models.User{}
	err := s.repository.Find(&users).Error
	if err != nil {
		return nil, httperror.NewHttpError(constants.ERROR_GETTING_USER, "", http.StatusInternalServerError)
	}
	return mapper.ToUsersResponseDto(users), err
}

// GetAllUser get all the user
func (s UserService) GetAllUser(searchQuery string, pagination utils.Pagination) (utils.Pagination, error) {

	query := s.repository.Table("users")

	// Filter search
	if searchQuery != "" {
		s.logger.Info("Search query: ", searchQuery)
		query = query.Where("users.name LIKE ?", "%"+searchQuery+"%").
			Or("users.email LIKE ?", "%"+searchQuery+"%")
	}

	user := []models.User{}
	err := query.Scopes(paginate(&user, &pagination, query)).
		Order(pagination.GetSort()).Find(&user).Error
	if err != nil {
		return pagination, httperror.NewHttpError(constants.ERROR_GETTING_USER, "", http.StatusInternalServerError)
	}

	pagination.Rows = mapper.ToUsersResponseDto(user)

	return pagination, err
}

// CreateUser call to create the user
func (s UserService) CreateUser(createUserDto dto.CreateUserRequest) error {
	// Generate password
	// password := utils.GenerateRandomString(12)
	// createUserDto.Password = password

	user, _ := mapper.ToUserModel(createUserDto)

	// sendEmailRequest := dto.SendEmailRequestDto{
	// 	To:      createUserDto.Email,
	// 	Subject: "Gello",
	// 	Body:    "Your password is: " + createUserDto.Password,
	// }

	// s.logger.Debug("SEND EMAIL: ", sendEmailRequest)

	// // Send password to user email
	// err := s.smtpClient.SendEmail(sendEmailRequest)
	// if err != nil {
	// 	return httperror.NewHttpError(constants.ERROR_SENDING_EMAIL, "", http.StatusInternalServerError)
	// }

	return s.repository.Create(&user).Error
}

// UpdateUser updates the user
func (s UserService) UpdateUser(id uint, dto dto.CreateUserRequest) error {
	user, _ := mapper.ToUserModel(dto)
	s.logger.Info("USER: ", user)
	err := s.repository.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return httperror.NewHttpError(constants.ERROR_UPDATING_USER, "", http.StatusInternalServerError)
	}
	return err
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	err := s.repository.Delete(&models.User{}, id).Error
	if err != nil {
		return httperror.NewHttpError(constants.ERROR_DELETING_USER, "", http.StatusInternalServerError)
	}
	return err
}

// GetOneByEmail gets one user by email
func (s UserService) GetOneByEmail(email string) (user *models.User, err error) {
	err = s.repository.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusNotFound)
	}

	return user, nil
}

// GetOneUserByEmail implements domains.UserService.
func (s UserService) GetOneUserByEmail(email string) (dto dto.UserResponseDto, err error) {
	user := models.User{}
	err = s.repository.First(&user, "email = ? ", email).Error
	if err != nil {
		return dto, httperror.NewHttpError(constants.ERROR_USER_NOT_FOUND, "", http.StatusNotFound)
	}
	return mapper.ToUserResponseDto(user), err
}
