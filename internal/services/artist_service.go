package services

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type ArtistService interface {
	CreateArtist(req models.ArtistCreateRequest) (*models.Artist, error)
	GetArtist(id uint) (*models.Artist, error)
	UpdateArtist(id uint, req models.ArtistUpdateRequest) (*models.Artist, error)
	DeleteArtist(id uint) error
	ListArtist() ([]models.Artist, error)
}

type artistService struct {
	artistRepo repository.ArtistRepository
}

func NewArtistService(
	artistRepo repository.ArtistRepository,
) ArtistService {
	return &artistService{
		artistRepo: artistRepo,
	}
}

func (s *artistService) CreateArtist(req models.ArtistCreateRequest) (*models.Artist, error) {

	artist := models.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	if err := s.artistRepo.Create(&artist); err != nil {
		return nil, err
	}

	return &artist, nil
}

func (s *artistService) GetArtist(id uint) (*models.Artist, error) {
	return s.artistRepo.GetByID(id)
}

func (s *artistService) UpdateArtist(id uint, req models.ArtistUpdateRequest) (*models.Artist, error) {
	artist, err := s.artistRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		artist.Name = *req.Name
	}

	if req.Bio != nil {
		artist.Bio = *req.Bio
	}

	if err := s.artistRepo.Update(artist); err != nil {
		return nil, err
	}
	return artist, nil
}

func (s *artistService) DeleteArtist(id uint) error {
	if _, err := s.artistRepo.GetByID(id); err != nil {
		return err
	}

	return s.artistRepo.Delete(id)
}

func (s *artistService) ListArtist() ([]models.Artist, error) {
	return s.artistRepo.List()
}
