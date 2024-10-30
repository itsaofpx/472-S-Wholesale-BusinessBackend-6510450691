package request

type FilterProductRequest struct {
	Name     string  `json:"name"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
}

type BuyProductRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
