package model

import "gorm.io/gorm"

type QuestionImage struct {
	gorm.Model
	QuestionID uint
	Key        string `gorm:"unique;not null"`
	URL        string `gorm:"not null"`
}
