package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string         `gorm:"size:256;not null"`
	Username      string         `gorm:"size:25;not null;unique"`
	Email         string         `gorm:"not null;unique"`
	Password      string         `gorm:"not null"`
	Questions     []Question     `gorm:"constraint:OnDelete:CASCADE"`
	Answers       []Answer       `gorm:"constraint:OnDelete:CASCADE"`
	Favorites     []*Question    `gorm:"many2many:favorites"`
	Ratings       []Rating       `gorm:"constraint:OnDelete:CASCADE"`
	Notifications []Notification `gorm:"constraint:OnDelete:CASCADE"`
	Memberships   []Membership   `gorm:"constraint:OnDelete:CASCADE"`
}
