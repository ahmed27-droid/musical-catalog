package repository

import (
	"musical-catalog/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error
	GetAll(trackID *uint) ([]models.Review, error)
	Delete(id uint) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) GetAll(trackID *uint) ([]models.Review, error) {
	var reviews []models.Review

	query := r.db
	if trackID != nil {
		query = query.Where("track_id = ?", *trackID)
	}

	err := query.Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}
