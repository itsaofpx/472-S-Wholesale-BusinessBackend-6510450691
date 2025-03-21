package entities

type Product struct {
	Id            int     `json:"id"`
	P_name        string  `json:"p_name"`
	P_description string  `json:"p_description"`
	P_location    string  `json:"p_location"`
	P_amount      int     `json:"p_amount"`
	P_price       float64 `json:"p_price"`
	Image_url_1   string  `json:"image_url_1"`
	Image_url_2   string  `json:"image_url_2"`
	Image_url_3   string  `json:"image_url_3"`
}
