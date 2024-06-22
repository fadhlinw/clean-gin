package repository

import (
	"github.com/fadhlinw/clean-gin/lib"
	"gorm.io/gorm"
)

// TokenStoreRepository
type TokenStoreRepository struct {
	lib.Database
	logger lib.Logger
}

// NewTokenStoreRepository creates a new token store repository
func NewTokenStoreRepository(db lib.Database, logger lib.Logger) TokenStoreRepository {
	return TokenStoreRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r TokenStoreRepository) WithTrx(trxHandle *gorm.DB) TokenStoreRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
