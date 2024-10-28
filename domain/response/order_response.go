package response

import (
	"time"
)

type OrderResponse struct {
	Id            int       `json:"id"`            // Set primary key with auto-increment
	O_status      string    `json:"o_status"`      // Not-null constraint
	O_timestamp   time.Time `json:"o_timestamp"`   // Not-null constraint
	O_total_price float64   `json:"o_total_price"` // Not-null constraint
	UserID        int       `json:"userID"`
}
