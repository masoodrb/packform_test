package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id         int64 `gorm:"primaryKey"`
	OrderName  string
	CustomerId string
}

type OrderItem struct {
	gorm.Model
	Id           int64 `gorm:"primaryKey"`
	OrderId      int64
	PricePerUnit float64
	Quantity     int64
	Product      string
}

type Delivery struct {
	gorm.Model
	Id                int64 `gorm:"primaryKey"`
	OrderItemId       int64
	DeliveredQuantity int64
}

type Customers struct {
	gorm.Model
	UserId      string `gorm:"primaryKey"`
	Login       string
	Password    string
	Name        string
	CompanyId   int64
	CreditCards string
}

type CustomerComapnies struct {
	gorm.Model
	CustomerId  int64  `gorm:"primaryKey"`
	CompanyName string `gorm:"primaryKey"`
}
