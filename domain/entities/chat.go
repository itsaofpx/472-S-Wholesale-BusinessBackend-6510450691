package entities

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	UserID   int       `gorm:"not null"`
	Messages []Message `gorm:"foreignKey:ChatID"`
}
