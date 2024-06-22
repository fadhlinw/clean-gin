// Package services provides services related to OTP (One-Time Password).

package services

import (
	"net/http"

	"github.com/fadhlinw/clean-gin/constants"
	"github.com/fadhlinw/clean-gin/domains"
	httperror "github.com/fadhlinw/clean-gin/error"
	"github.com/fadhlinw/clean-gin/lib"
	"github.com/fadhlinw/clean-gin/models"
	"github.com/fadhlinw/clean-gin/repository"
)

// OTPService service relating to otp
type OTPService struct {
	env        lib.Env
	logger     lib.Logger
	repository repository.OTPRepository
}

// NewOTPService creates a new otp service
func NewOTPService(env lib.Env, logger lib.Logger, repository repository.OTPRepository) domains.OTPService {
	return OTPService{
		env:        env,
		logger:     logger,
		repository: repository,
	}
}

// Create saves the OTP data to the database.
func (s OTPService) Create(userId int, code string) error {
	var otp = models.Otp{
		UserId: userId,
		Code:   code,
		IsUsed: false,
	}
	err := s.repository.Create(&otp).Error

	if err != nil {
		s.logger.Error("Error saving OTP data to DB")
		s.logger.Debug("Detail: ", err.Error())
		return httperror.NewHttpError(constants.ERROR_CREATING_OTP, "", http.StatusInternalServerError)
	}
	return nil
}

// UpdateById updates the OTP data in the database by ID.
func (s OTPService) UpdateById(id int, isUsed bool) error {
	s.logger.Info("Updating OTP data to DB")
	err := s.repository.Model(&models.Otp{}).Where("id = ?", id).Update("is_used", true).Error

	if err != nil {
		s.logger.Error("Error updating OTP data to DB")
		s.logger.Debug("Detail: ", err.Error())
		return httperror.NewHttpError(constants.ERROR_UPDATING_OTP, "", http.StatusInternalServerError)
	}
	return nil
}

// GetByCode retrieves the OTP data from the database by code.
func (s OTPService) GetByCode(userId int, code string) (*models.Otp, error) {
	s.logger.Info("Getting OTP data from DB")
	var otp models.Otp
	err := s.repository.Where("user_id = ? AND code = ? AND is_used = false", userId, code).First(&otp).Error

	if err != nil {
		s.logger.Error("Error getting OTP data from DB")
		s.logger.Debug("Detail: ", err.Error())
		return nil, httperror.NewHttpError(constants.ERROR_GETTING_OTP_BY_CODE, "", http.StatusInternalServerError)
	}
	return &otp, nil
}

// GetByUserIdAndIsUsed retrieves the OTP data from the database by user ID and isUsed flag.
func (s OTPService) GetByUserIdAndIsUsed(userId int, isUsed bool) (*models.Otp, error) {
	s.logger.Info("Getting OTP data from DB")
	var otp models.Otp
	err := s.repository.Where("user_id = ? AND is_used = ?", userId, isUsed).First(&otp).Error

	if err != nil {
		s.logger.Error("Error getting OTP data from DB")
		s.logger.Debug("Detail: ", err.Error())
		return nil, err
	}
	return &otp, nil
}
