package repository

import (
	"musical-catalog/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	List() ([]models.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Playlists").Preload("Reviews").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *gormUserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *gormUserRepository) List() ([]models.User, error) {
	var user []models.User

	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
