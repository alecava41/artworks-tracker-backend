package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ArtworkService service layer
type ArtworkService struct {
	logger     lib.Logger
	repository repository.ArtworkRepository
}

// NewArtworkService creates a new ArtworkService
func NewArtworkService(logger lib.Logger, repository repository.ArtworkRepository) ArtworkService {
	return ArtworkService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s ArtworkService) WithTrx(trxHandle *gorm.DB) ArtworkService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneArtwork gets one user
func (s ArtworkService) GetOneArtwork(beaconId uuid.UUID, lan string) (artwork models.Artwork, err error) {
	return artwork,
		s.repository.
			Raw(`SELECT author, title, description, beacon
					 FROM artworks a INNER JOIN translations t ON a.id = t.artwork 
					 WHERE language = ? AND a.beacon = ?`, lan, beaconId).
			First(&artwork).Error
}
