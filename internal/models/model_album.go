package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model

	Title string `json:"title" gorm:"not null;size:20;uniqueIndex:idx_artist_title"`
	Year  int    `json:"year" gorm:"not null"`

	ArtistID uint   `json:"artist_id" gorm:"not null;uniqueIndex:idx_artist_title"`
	Artist   Artist `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Tracks []Track `json:"tracks,omitempty" gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
}

type AlbumCreateRequest struct {
	Title string `json:"title" binding:"required,min=2,max=20"`
	Year  int    `json:"year" binding:"required,gte=1950,lte=2025"`

	ArtistID uint `json:"artist_id" binding:"required"`
}

type AlbumUpdateRequest struct {
	ArtistID *uint   `json:"artist_id" binding:"omitempty"`
	Title    *string `json:"title" binding:"omitempty,min=2,max=20"`
	Year     *int    `json:"year" binding:"omitempty,gte=1950,lte=2025"`
}
