package models

import (
	"time"

	"gorm.io/gorm"
)

type NotificationType string

const (
	NewRating   NotificationType = "rating"
	BestAnswer  NotificationType = "best-answer"
	NewAnswer   NotificationType = "new-answer"
	NewFavorite NotificationType = "favorite"
)

type Notification struct {
	gorm.Model
	UserID            uint
	QuestionID        uint
	Type              NotificationType `gorm:"not null"`
	TransmitterUserID uint             `gorm:"not null"`
	Description       string           `gorm:"not null"`
	Rating            uint
	ReadAt            time.Time
}
