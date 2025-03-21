package entities

type Supplier struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	Contract_person string `json:"contract_person"`
	Email           string `json:"email	"`
}
