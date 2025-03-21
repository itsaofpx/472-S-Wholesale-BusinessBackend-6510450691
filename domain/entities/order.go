package entities

import (
	"time"
)

type Order struct {
	Id             int         `gorm:"primaryKey" json:"id"`              // Set primary key with auto-increment
	O_status       string      `gorm:"not null" json:"o_status"`          // Not-null constraint
	O_timestamp    time.Time   `gorm:"not null" json:"o_timestamp"`       // Not-null constraint
	O_total_price  float64     `gorm:"not null" json:"o_total_price"`     // Not-null constraint
	UserID         int         `gorm:"not null" json:"userID"`            // Foreign key with not-null
	User           User        `gorm:"foreignKey:UserID"`  
	TrackingNumber string      `json:"tracking_number"`                 // Foreign key association
	
}
