package repository

import (
	"github.com/fadhlinw/clean-gin/lib"
	"gorm.io/gorm"
)

// OTPRepository database structure
type OTPRepository struct {
	lib.Database
	logger lib.Logger
}

// NewOTPRepository creates a new otp repository
func NewOTPRepository(db lib.Database, logger lib.Logger) OTPRepository {
	return OTPRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r OTPRepository) WithTrx(trxHandle *gorm.DB) OTPRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
