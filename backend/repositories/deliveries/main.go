package deliveries

import (
	"encoding/json"
	"fmt"
	"masoodrb/packform/utils"
	"time"
)

type OrderDetails struct {
	OrderName    string
	ProductName  string
	CompanyName  string
	CustomerName string
	OrderDate    time.Time
	Quantity     int64
	PricePerUnit float64
}

func GetAllDeliveryDetails() []OrderDetails {
	db := utils.GetDBContext()
	orderDetails := []OrderDetails{}

	db.Raw(`
			SELECT
				orders.order_name as order_name,
				order_items.product as product_name,
				companies.company_name as company_name,
				customers.name as customer_name,
				orders.created_at as order_date,
				deliveries.delivered_quantity as quantity,
				order_items.price_per_unit as price_per_unit
			FROM
				deliveries
				join order_items on order_items.id = deliveries.order_item_id
				join orders on order_items.order_id = orders.id
				join customers on orders.customer_id = customers.user_id
				join companies on customers.company_id = companies.id
		`).
		Find(&orderDetails)

	return orderDetails
}

func GetDeliveryDetails(searchString string) []OrderDetails {
	db := utils.GetDBContext()

	searchString = "%" + searchString + "%"

	orderDetails := []OrderDetails{}

	db.Raw(`
			SELECT
				orders.order_name as order_name,
				order_items.product as product_name,
				companies.company_name as company_name,
				customers.name as customer_name,
				orders.created_at as order_date,
				deliveries.delivered_quantity as quantity,
				order_items.price_per_unit as price_per_unit
			FROM
				deliveries
				join order_items on order_items.id = deliveries.order_item_id
				join orders on order_items.order_id = orders.id
				join customers on orders.customer_id = customers.user_id
				join companies on customers.company_id = companies.id
			WHERE
				orders.order_name like ? or order_items.product like ? 
		`, searchString, searchString).
		Find(&orderDetails)

	return orderDetails
}

func GetDeliveryDetailsByDate(startDate time.Time, endDate time.Time, searchString string) []OrderDetails {
	db := utils.GetDBContext()
	orderDetails := []OrderDetails{}
	searchString = "%" + searchString + "%"
	sql := `
		SELECT
			orders.order_name as order_name,
			order_items.product as product_name,
			companies.company_name as company_name,
			customers.name as customer_name,
			orders.created_at as order_date,
			deliveries.delivered_quantity as quantity,
			order_items.price_per_unit as price_per_unit
		FROM
			deliveries
			join order_items on order_items.id = deliveries.order_item_id
			join orders on order_items.order_id = orders.id
			join customers on orders.customer_id = customers.user_id
			join companies on customers.company_id = companies.id
		WHERE
			orders.created_at between ? and ?
	`
	if searchString != "" {
		sql = sql + `
			AND (orders.order_name like ? or order_items.product like ? )
		`
		db.Raw(sql, startDate, endDate, searchString, searchString).
			Find(&orderDetails)
	} else {
		db.Raw(sql, startDate, endDate).
			Find(&orderDetails)
	}

	return orderDetails
}

func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}
