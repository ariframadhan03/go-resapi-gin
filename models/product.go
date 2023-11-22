package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"varchar(300)" json:"product_name"`
	Description string `gorm:"text" json:"description"`
}
