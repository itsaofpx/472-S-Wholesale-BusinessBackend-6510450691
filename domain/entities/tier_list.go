package entities

type TierList struct {
	Tier            int     `json:"tier" gorm:"primaryKey"` // 'Tier' is the primary key
	DiscountPercent float64 `json:"discount_percent"`
}
