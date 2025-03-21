package response

type OrderLineResponse struct {
	ID           int     `json:"id"`
	OrderID      int     `json:"order_id"`
	ProductID    int     `json:"product_id"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	ProductImg   string  `json:"product_img"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductAmount int    `json:"product_amount"`
}