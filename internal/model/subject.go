package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string `gorm:"size:50;not null"`
	// Questions []Question `gorm:"constraint:OnDelete:CASCADE"`
}
