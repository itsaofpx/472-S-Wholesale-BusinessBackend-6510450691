package request

type UpdateOrderStatusRequest struct {
	ID int `json:"id"`
	Status string `json:"o_status"`
}