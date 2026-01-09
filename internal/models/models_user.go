package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `json:"name" gorm:"not null;size:20;uniqueIndex"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null;size:50"`

	Reviews   []Review   `gorm:"constraint:OnDelete:CASCADE;"`
	Playlists []Playlist `gorm:"constraint:OnDelete:CASCADE;"`
}

type CreateUser struct {
	Name  string `json:"name" binding:"required,min=5,max=20"`
	Email string `json:"email" binding:"required,max=50,email"`
}

type UpdateUser struct {
	Name  *string `json:"name" binding:"min=5,max=20"`
	Email *string `json:"email" binding:"max=50,email"`
}

