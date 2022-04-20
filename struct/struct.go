package structs

import "time"

type Orders struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Items   `gorm:"foreignKey:OrderId"`
}

type Items struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"orderId"`
}

type CreateOrders struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Items   `json:"items"`
}
