package services

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
)

type UserService interface {
	CreateUser(req models.CreateUser) (*models.User, error)
	GetByIDUser(id uint) (*models.User, error)
	UpdateUser(id uint, req models.UpdateUser) (*models.User, error)
	DeleteUser(id uint) error
	ListUser() ([]models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(req models.CreateUser) (*models.User, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetByIDUser(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) UpdateUser(id uint, req models.UpdateUser) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return  err
	}
	return s.userRepo.Delete(id)
}

func (s *userService) ListUser() ([]models.User, error){
	return s.userRepo.List()
}