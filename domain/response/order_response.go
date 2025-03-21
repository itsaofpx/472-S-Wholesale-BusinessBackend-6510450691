package response

import (
	"time"
)

type OrderResponse struct {
	Id             int       `json:"id"`            // Set primary key with auto-increment
	O_status       string    `json:"o_status"`      // Not-null constraint
	O_timestamp    time.Time `json:"o_timestamp"`   // Not-null constraint
	O_total_price  float64   `json:"o_total_price"` // Not-null constraint
	UserID         int       `json:"userID"`
	TrackingNumber string  `json:"tracking_number"`
}

type OrderUserResponse struct {
	Id             int       `json:"id"`            // Set primary key with auto-increment
	O_status       string    `json:"o_status"`      // Not-null constraint
	O_timestamp    time.Time `json:"o_timestamp"`   // Not-null constraint
	O_total_price  float64   `json:"o_total_price"` // Not-null constraint
	UserID         int       `json:"userID"`
	TrackingNumber string    `json:"tracking_number"`
	F_name         string    `json:"f_name"`
	L_name         string    `json:"l_name"`
	Phone_number   string    `json:"phone_number"`
	Email          string    `json:"email"`
	Tier_rank      int       `json:"tier_rank"`
	Address        string    `json:"address"`
}
