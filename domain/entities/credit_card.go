package entities

type CreditCard struct {
	ID             int    `json:"id" gorm:"primaryKey"`      // รหัสบัตรเครดิต (Primary Key)
	UserID         int 	`json:"user_id" gorm:"index;not null"` // อีเมลของเจ้าของบัตร ใช้เป็น Foreign Key เชื่อมกับ User
	CardNumber     string `json:"card_number"`               // หมายเลขบัตรเครดิต
	CardHolder     string `json:"card_holder"`               // ชื่อผู้ถือบัตร
	Expiration     string `json:"expiration"`                // วันหมดอายุของบัตร (MM/YY หรือ MM/YYYY)
	SecurityCode   string `json:"security_code"`             // รหัส CVV/CVC ของบัตร
}
