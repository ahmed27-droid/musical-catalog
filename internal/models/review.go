package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model

	UserID  uint `json:"user_id" gorm:"not null;uniqueIndex:idx_user_track"`
	TrackID uint `json:"track_id" gorm:"not null;uniqueIndex:idx_user_track"`

	Rating int    `json:"rating" gorm:"not null"`
	Text   string `json:"text" gorm:"type:text"`

	User  User  `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Track Track `json:"-" gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
}

type CreateReviewRequest struct {
	UserID  uint   `json:"user_id" binding:"required"`
	TrackID uint   `json:"track_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Text    string `json:"text"`
}
