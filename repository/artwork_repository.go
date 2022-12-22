package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// ArtworkRepository database structure
type ArtworkRepository struct {
	lib.Database
	logger lib.Logger
}

// NewArtworkRepository creates a new user repository
func NewArtworkRepository(db lib.Database, logger lib.Logger) ArtworkRepository {
	return ArtworkRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r ArtworkRepository) WithTrx(trxHandle *gorm.DB) ArtworkRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
