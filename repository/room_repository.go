package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// RoomRepository database structure
type RoomRepository struct {
	lib.Database
	logger lib.Logger
}

// NewRoomRepository creates a new user repository
func NewRoomRepository(db lib.Database, logger lib.Logger) RoomRepository {
	return RoomRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r RoomRepository) WithTrx(trxHandle *gorm.DB) RoomRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
