package services

import (
	"errors"
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type AlbumService interface {
	CreateAlbum(req models.AlbumCreateRequest) (*models.Album, error)
	GetAlbum(id uint) (*models.Album, error)
	UpdateAlbum(id uint, req models.AlbumUpdateRequest) (*models.Album, error)
	DeleteAlbum(id uint) error
	ListAlbums() ([]models.Album, error)

	GetAlbumsAverageRating(id uint) (float64, error)
}

type albumService struct {
	artistRepo repository.ArtistRepository
	albumRepo  repository.AlbumRepository
}

func NewAlbumService(
	artistRepo repository.ArtistRepository,
	albumRepo repository.AlbumRepository,
) AlbumService {
	return &albumService{
		artistRepo: artistRepo,
		albumRepo:  albumRepo,
	}
}

func (s *albumService) CreateAlbum(req models.AlbumCreateRequest) (*models.Album, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	album := models.Album{
		Title:    req.Title,
		Year:     req.Year,
		ArtistID: req.ArtistID,
	}

	if err := s.albumRepo.Create(&album); err != nil {
		return nil, err
	}

	return &album, nil
}

func (s *albumService) GetAlbum(id uint) (*models.Album, error) {
	return s.albumRepo.GetByID(id)
}

func (s *albumService) UpdateAlbum(id uint, req models.AlbumUpdateRequest) (*models.Album, error) {
	album, err := s.albumRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		album.Title = *req.Title
	}

	if req.Year != nil {
		album.Year = *req.Year
	}

	if req.ArtistID != nil {
		if _, err := s.artistRepo.GetByID(*req.ArtistID); err != nil {
			return nil, errors.New("artist not found")
		}
		album.ArtistID = *req.ArtistID
	}

	if err := s.albumRepo.Update(album); err != nil {
		return nil, err
	}
	return album, nil
}

func (s *albumService) DeleteAlbum(id uint) error {
	if _, err := s.albumRepo.GetByID(id); err != nil {
		return err
	}

	return s.albumRepo.Delete(id)
}

func (s *albumService) ListAlbums() ([]models.Album, error) {
	return s.albumRepo.List()
}

func (s *albumService) GetAlbumsAverageRating(id uint) (float64, error) {
	if id == 0 {
		return 0, errors.New("invalid album id")
	}

	if _, err := s.albumRepo.GetByID(id); err != nil {
		return 0, err
	}

	return s.albumRepo.GetAverageRating(id)
}
