package entities

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Timestamp time.Time `gorm:"not null"`
	Body      string    `gorm:"not null"`
	UserID    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	ChatID    int       `gorm:"not null"`
}
