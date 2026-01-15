package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null;size:20"`
	UserID uint   `json:"user_id" gorm:"not null;index"`

	User   *User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tracks []Track `json:"tracks" gorm:"many2many:playlist_tracks;constraint:OnDelete:CASCADE;"`
}

type CreatePlaylist struct {
	Title  string `json:"title" binding:"required,min=2,max=20"`
	UserID uint   `json:"user_id" binding:"required"`
}

type UpdatePlaylist struct {
	Title *string `json:"title" binding:"min=2,max=20"`
}
