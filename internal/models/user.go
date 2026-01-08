package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null;size:20;uniqueIndex"`
	Email string `json:"email" gorm:"uniqueIndex;not null;size:50"`
}

type CreateUser struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUser struct{
	Name string `json:"name" binding:"max=20"`
	Email string `json:"email" binding:"max=50"`
}