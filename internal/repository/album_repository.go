package repository

import (
	"musical-catalog/internal/models"

	"gorm.io/gorm"
)

type AlbumRepository interface {
	Create(album *models.Album) error
	GetByID(id uint) (*models.Album, error)
	Update(album *models.Album) error
	Delete(id uint) error
	List() ([]models.Album, error)
	GetAverageRating(albumID uint) (float64, error)
}

type gormAlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &gormAlbumRepository{db: db}
}

func (r *gormAlbumRepository) Create(album *models.Album) error {
	return r.db.Create(album).Error
}

func (r *gormAlbumRepository) GetByID(id uint) (*models.Album, error) {
	var album models.Album

	if err := r.db.Preload("Tracks").Preload("Artist").First(&album, id).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

func (r *gormAlbumRepository) Update(album *models.Album) error {
	return r.db.Save(album).Error
}

func (r *gormAlbumRepository) Delete(id uint) error {
	return r.db.Delete(&models.Album{}, id).Error
}

func (r *gormAlbumRepository) List() ([]models.Album, error) {
	var albums []models.Album

	if err := r.db.Preload("Artist").Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *gormAlbumRepository) GetAverageRating(albumID uint) (float64, error) {
	var avg float64

	err:= r.db.
		Model(&models.Track{}).
		Where("album_id = ?", albumID).
		Select("COALESCE(AVG(rating), 0)").
		Scan(&avg).Error


	return avg, err

}
