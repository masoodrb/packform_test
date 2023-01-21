package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id         int64 `gorm:"primaryKey"`
	OrderName  string
	CustomerId string
	Customer   Customer `gorm:"references:UserId"`
}

type OrderItem struct {
	Id           int64 `gorm:"primaryKey"`
	OrderId      int64
	Order        Order
	PricePerUnit float64
	Quantity     int64
	Product      string
}

type Delivery struct {
	Id                int64 `gorm:"primaryKey"`
	OrderItemId       int64
	OrderItem         OrderItem
	DeliveredQuantity int64
}

type Customer struct {
	UserId      string `gorm:"primaryKey"`
	Login       string
	Password    string
	Name        string
	CompanyId   int64
	Company     Company
	CreditCards string
}

type Company struct {
	Id          int64 `gorm:"primaryKey"`
	CompanyName string
}
