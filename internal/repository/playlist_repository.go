package repository

import (
	"musical-catalog/internal/models"

	"gorm.io/gorm"
)

type PlaylistRepository interface {
	Create(playlist *models.Playlist) error
	GetByID(id uint) (*models.Playlist, error)
	Delete(id uint) error
	AddTrack(playlistID uint, trackID uint) error
	DeleteTrack(playlistID uint, trackID uint) error
	List() ([]models.Playlist, error)
}

type gormPlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) PlaylistRepository {
	return &gormPlaylistRepository{db: db}
}

func (r *gormPlaylistRepository) Create(playlist *models.Playlist) error {
	return r.db.Create(playlist).Error
}

func (r *gormPlaylistRepository) GetByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist

	if err := r.db.Preload("Tracks").First(&playlist, id).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *gormPlaylistRepository) Delete(id uint) error {
	return r.db.Delete(&models.Playlist{}, id).Error
}

func (r *gormPlaylistRepository) AddTrack(playlistID uint, trackID uint) error {
	var playlist models.Playlist
	if err := r.db.First(&playlist, playlistID).Error; err != nil {
		return err
	}

	var track Track
	if err := r.db.First(&track, trackID).Error; err != nil {
		return err
	}

	return r.db.Model(&playlist).Association("Tracks").Append(&track)
}

func (r *gormPlaylistRepository) DeleteTrack(playlistID uint, trackID uint) error {
	var playlist models.Playlist
	if err := r.db.First(&playlist, playlistID).Error; err != nil {
		return err
	}

	var track Track
	if err := r.db.First(&track, trackID).Error; err != nil {
		return err
	}

	return r.db.Model(&playlist).Association("Tracks").Delete(&track)
}

func (r *gormPlaylistRepository) List() ([]models.Playlist, error) {
	var playlist []models.Playlist

	if err := r.db.Find(&playlist).Error; err != nil {
		return nil, err
	}
	return playlist, nil
}
