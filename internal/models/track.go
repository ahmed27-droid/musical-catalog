package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null;size:100"`
	Duration int    `json:"duration" gorm:"not null"` 

	AlbumID uint  `json:"album_id" gorm:"not null;index"` 
	Album   Album `json:"-" gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
}

type CreateTrackRequest struct {
	Title    string `json:"title" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
	AlbumID  uint   `json:"album_id" binding:"required"`
}

type UpdateTrackRequest struct {
	Title    *string `json:"title,omitempty"`
	Duration *int    `json:"duration,omitempty"`
	AlbumID  *uint   `json:"album_id,omitempty"`
}
