package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model

	Name   string  `json:"name" gorm:"not null;size:20;uniqueIndex"`
	Bio    string  `json:"bio" gorm:"not null;size:100"`
	Albums []Album `json:"albums,omitempty" gorm:"foreignKey:ArtistID"`
}

type ArtistCreateRequest struct {
	Name string `json:"name" binding:"required,min=2,max=20"`
	Bio  string `json:"bio" binding:"required,min=10,max=100"`
}

type ArtistUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty,min=2,max=20"`
	Bio  *string `json:"bio" binding:"omitempty,min=10,max=100"`
}
