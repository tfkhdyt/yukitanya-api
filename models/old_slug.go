package models

import "gorm.io/gorm"

type OldSlug struct {
	gorm.Model
	QuestionID uint
	Slug       string `gorm:"size:50;not null;unique"`
}
