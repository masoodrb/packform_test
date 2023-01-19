package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id         uint `gorm:"primaryKey"`
	OrderName  sql.NullString
	CustomerId sql.NullString
}

type OrderItem struct {
	gorm.Model
	Id           uint `gorm:"primaryKey"`
	OrderId      int
	PricePerUnit float32
	Quantity     int64
	Product      sql.NullString
}

type Delivery struct {
	gorm.Model
	Id                uint `gorm:"primaryKey"`
	OrderItemId       uint
	DeliveredQuantity uint
}

type Customers struct {
	gorm.Model
	UserId      sql.NullString `gorm:"primaryKey"`
	Login       sql.NullString
	Password    sql.NullString
	Name        sql.NullString
	CompanyId   uint
	CreditCards string
}

type CustomerComapnies struct {
	gorm.Model
	CustomerId  uint           `gorm:"primaryKey"`
	CompanyName sql.NullString `gorm:"primaryKey"`
}
