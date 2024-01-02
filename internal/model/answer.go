package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	QuestionID   uint
	UserID       uint
	Content      string   `gorm:"size:1000;not null"`
	IsBestAnswer string   `gorm:"not null;default:false"`
	Ratings      []Rating `gorm:"constraint:OnDelete:CASCADE"`
}
