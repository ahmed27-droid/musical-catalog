package services

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type PlaylistService interface {
	CreatePlaylist(req models.CreatePlaylist) (*models.Playlist, error)
	GetByIDPlaylist(id uint) (*models.Playlist, error)
	DeletePlaylist(id uint) error
	AddTrackToPlaylist(playlistID uint, trackID uint) error
	DeleteTrackOfPlaylist(playlistID uint, trackID uint) error
	ListPlaylist() ([]models.Playlist, error)
}

type playlistService struct {
	userRepo     repository.UserRepository
	playlistRepo repository.PlaylistRepository
	trackRepo repository.TrackRepository
}

func NewPlaylistService(
	userRepo repository.UserRepository,
	playlistRepo repository.PlaylistRepository,
	trackRepo repository.TrackRepository,
) PlaylistService {
	return &playlistService{
		userRepo:     userRepo,
		playlistRepo: playlistRepo,
		trackRepo:    trackRepo,
	}
}

func (s *playlistService) CreatePlaylist(req models.CreatePlaylist) (*models.Playlist, error) {

	playlist := models.Playlist{
		Title:  req.Title,
		UserID: req.UserID,
	}

	if err := s.playlistRepo.Create(&playlist); err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (s *playlistService) GetByIDPlaylist(id uint) (*models.Playlist, error) {
	return s.playlistRepo.GetByID(id)
}
func (s *playlistService) DeletePlaylist(id uint) error {
	_, err := s.playlistRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.playlistRepo.Delete(id)
}

func (s *playlistService) AddTrackToPlaylist(playlistID uint, trackID uint) error {
	if _, err := s.playlistRepo.GetByID(playlistID); err != nil {
		return err
	}

	if _, err := s.trackRepo.GetByID(trackID); err != nil {
		return err
	}

	return s.playlistRepo.AddTrack(playlistID, trackID)
}

func (s *playlistService) DeleteTrackOfPlaylist(playlistID uint, trackID uint) error {
	if _, err := s.playlistRepo.GetByID(playlistID); err != nil {
		return err
	}

	if _, err := s.trackRepo.GetByID(trackID); err != nil {
		return err
	}

	return s.playlistRepo.DeleteTrack(playlistID, trackID)
}

func (s *playlistService) ListPlaylist() ([]models.Playlist, error) {
	return s.playlistRepo.List()
}
