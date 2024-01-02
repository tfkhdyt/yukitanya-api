package model

import (
	"time"

	"gorm.io/gorm"
)

type MembershipType string

const (
	Standard MembershipType = "standard"
	Plus     MembershipType = "Plus"
)

type Membership struct {
	gorm.Model
	UserID    uint
	Type      MembershipType `gorm:"size:10;not null"`
	ExpiresAt time.Time      `gorm:"not null"`
}
