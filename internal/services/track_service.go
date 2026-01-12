package services

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type TrackService interface {
	Create(req models.CreateTrackRequest) (*models.Track, error)
	GetAll() ([]models.Track, error)
	GetByID(id uint) (*models.Track, error)
	Update(id uint, req models.UpdateTrackRequest) (*models.Track, error)
	Delete(id uint) error
	GetAverage(trackID uint) (float64, error)
}

type trackService struct {
	trackRepo repository.TrackRepository
}

func NewTrackService(trackRepo repository.TrackRepository) TrackService {
	return &trackService{trackRepo: trackRepo}
}

func (s *trackService) Create(req models.CreateTrackRequest) (*models.Track, error) {
	track := &models.Track{
		Title:    req.Title,
		Duration: req.Duration,
		AlbumID:  req.AlbumID,
	}

	if err := s.trackRepo.Create(track); err != nil {
		return nil, err
	}

	return track, nil
}

func (s *trackService) GetAll() ([]models.Track, error) {
	return s.trackRepo.GetAll()
}

func (s *trackService) GetByID(id uint) (*models.Track, error) {
	return s.trackRepo.GetByID(id)
}

func (s *trackService) Update(id uint, req models.UpdateTrackRequest) (*models.Track, error) {
	track, err := s.trackRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		track.Title = *req.Title
	}
	if req.Duration != nil {
		track.Duration = *req.Duration
	}
	if req.AlbumID != nil {
		track.AlbumID = *req.AlbumID
	}

	if err := s.trackRepo.Update(track); err != nil {
		return nil, err
	}

	return track, nil
}

func (s *trackService) Delete(id uint) error {
	return s.trackRepo.Delete(id)
}

func (s *trackService) GetAverage(trackID uint) (float64, error) {
	return s.trackRepo.GetAverageRating(trackID)
}
