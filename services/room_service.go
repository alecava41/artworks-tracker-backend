package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoomService service layer
type RoomService struct {
	logger     lib.Logger
	repository repository.RoomRepository
}

// NewRoomService creates a new RoomService
func NewRoomService(logger lib.Logger, repository repository.RoomRepository) RoomService {
	return RoomService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s RoomService) WithTrx(trxHandle *gorm.DB) RoomService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneRoom gets one room
func (s RoomService) GetOneRoom(beaconId uuid.UUID, lan string) (room models.Room, err error) {

	// Get the room id
	var tinyArtwork = new(models.ArtworkInfo)
	s.repository.Find(&tinyArtwork, "beacon = ?", beaconId)

	// Get all the ArtworkInfo
	var artworksInfo = make([]*models.ArtworkInfo, 0)
	s.repository.Raw(`
		SELECT * 
		FROM artworks a INNER JOIN translations t ON a.id = t.artwork
		WHERE room = ? and language = ?
	`, tinyArtwork.RoomId, lan).Find(&artworksInfo)

	err = s.repository.First(&room, tinyArtwork.RoomId).Error
	room.Artworks = artworksInfo

	return room, err
}
