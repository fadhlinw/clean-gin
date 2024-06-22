package domains

import (
	"github.com/fadhlinw/clean-gin/dto"
	"github.com/fadhlinw/clean-gin/models"
	"gorm.io/gorm"
)

type AuthService interface {
	WithTrx(trxHandle *gorm.DB) AuthService
	Authorize(tokenString string, claimType string) (*dto.AuthIdentityDto, error)
	CreateToken(user *models.User, claimType string) (string, error)
	CreateRefreshToken(user models.User) (string, error)
	ValidateAuth(request *dto.AuthRequestDto, refreshTokenString string) (*dto.AuthResponseDto, error)
	SaveToken(token string) error
	ChangePassword(id int, request *dto.AuthChangePasswordDto) error
	ForgotPassword(request dto.AuthForgotPasswordDto) error
	ValidateOTP(request dto.ValidateOTPRequestDto) (*dto.ValidateOTPResponseDto, error)
	ResetPassword(request dto.AuthResetPasswordDto) error
	Logout(tokenString string, refreshToken string) error
}
