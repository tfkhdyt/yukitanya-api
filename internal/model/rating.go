package model

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	AnswerID uint
	UserID   uint
	Value    uint `gorm:"size:2;not null"`
}
