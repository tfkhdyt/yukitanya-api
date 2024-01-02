package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	SubjectID uint
	UserID    uint
	Content   string          `gorm:"size:1000;not null"`
	Slug      string          `gorm:"size:50;not null;unique"`
	OldSlugs  []OldSlug       `gorm:"constraint:OnDelete:CASCADE"`
	Answers   []Answer        `gorm:"constraint:OnDelete:CASCADE"`
	Images    []QuestionImage `gorm:"constraint:OnDelete:CASCADE"`
	Favorites []*User         `gorm:"many2many:favorites"`
}
