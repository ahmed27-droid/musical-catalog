package repository

import (
	"musical-catalog/internal/models"
    "gorm.io/gorm"

)


type ArtistRepository interface{
	Create(artist *models.Artist) error
	GetByID(id uint) (*models.Artist, error)
	Update(artist *models.Artist) error
	Delete(id uint)error
	List()([]models.Artist,error)
}

type gormArtistRepository struct{
	db *gorm.DB
}


func NewArtistRepository(db *gorm.DB) ArtistRepository{
	return &gormArtistRepository{db:db}
}

func (r *gormArtistRepository) Create(artist *models.Artist) error{
	return r.db.Create(artist).Error
}

func (r *gormArtistRepository) GetByID(id uint) (*models.Artist, error){
	var artist models.Artist

	if err:= r.db.Preload("Albums").First(&artist, id).Error; err != nil{
		return nil, err
	}
	
	return &artist, nil
}

func (r *gormArtistRepository) Update(artist *models.Artist) error{
	return r.db.Save(artist).Error
}

func (r *gormArtistRepository) Delete(id uint) error{
	return r.db.Delete(&models.Artist{}, id).Error
}

func (r *gormArtistRepository) List() ([]models.Artist, error){
var artists []models.Artist

if err := r.db.Find(&artists).Error; err !=nil {
	return nil, err
}
	return artists, nil
}
