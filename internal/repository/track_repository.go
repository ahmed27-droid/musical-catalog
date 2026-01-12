package repository

import (
	"musical-catalog/internal/models"

	"gorm.io/gorm"
)

type TrackRepository interface {
	Create(track *models.Track) error
	GetAll() ([]models.Track, error)
	GetByID(id uint) (*models.Track, error)
	Update(track *models.Track) error
	Delete(id uint) error
	GetAverageRating(trackID uint) (float64, error)
}

type trackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) TrackRepository {
	return &trackRepository{db: db}
}

func (r *trackRepository) Create(track *models.Track) error {
	return r.db.Create(track).Error
}

func (r *trackRepository) GetAll() ([]models.Track, error) {
	var tracks []models.Track
	err := r.db.Find(&tracks).Error
	return tracks, err
}

func (r *trackRepository) GetByID(id uint) (*models.Track, error) {
	var track models.Track
	if err := r.db.First(&track, id).Error; err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *trackRepository) Update(track *models.Track) error {
	return r.db.Save(track).Error
}

func (r *trackRepository) Delete(id uint) error {
	return r.db.Delete(&models.Track{}, id).Error
}

func (r *trackRepository) GetAverageRating(trackID uint) (float64, error) {
	var avg float64
	err := r.db.
		Table("reviews").
		Select("AVG(rating)").
		Where("track_id = ?", trackID).
		Scan(&avg).Error

	return avg, err
}
