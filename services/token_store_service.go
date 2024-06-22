package services

import (
	"github.com/fadhlinw/clean-gin/domains"
	"github.com/fadhlinw/clean-gin/lib"
	"github.com/fadhlinw/clean-gin/models"
	"github.com/fadhlinw/clean-gin/repository"
	"gorm.io/gorm"
)

type TokenStoreService struct {
	logger     lib.Logger
	repository repository.TokenStoreRepository
}

func NewTokenStoreService(logger lib.Logger, repository repository.TokenStoreRepository) domains.TokenStoreService {
	return TokenStoreService{
		logger:     logger,
		repository: repository,
	}
}

func (s TokenStoreService) WithTrx(trxHandle *gorm.DB) domains.TokenStoreService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s TokenStoreService) CreateToken(tokenStore models.TokenStore) error {
	return s.repository.Create(&tokenStore).Error
}

func (s TokenStoreService) DeleteToken(token string) error {
	return s.repository.Where("token = ?", token).Delete(&models.TokenStore{}).Error
}

func (s TokenStoreService) ValidateToken(token string) error {
	return s.repository.Where("token = ?", token).First(&models.TokenStore{}).Error
}
