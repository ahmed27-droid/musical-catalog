package services

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type ReviewService interface {
	Create(req models.CreateReviewRequest) error
	GetAll(trackID *uint) ([]models.Review, error)
	Delete(id uint) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) Create(req models.CreateReviewRequest) error {
	review := models.Review{
		UserID:  req.UserID,
		TrackID: req.TrackID,
		Rating:  req.Rating,
		Text:    req.Text,
	}

	return s.repo.Create(&review)
}

func (s *reviewService) GetAll(trackID *uint) ([]models.Review, error) {
	return s.repo.GetAll(trackID)
}

func (s *reviewService) Delete(id uint) error {
	return s.repo.Delete(id)
}
