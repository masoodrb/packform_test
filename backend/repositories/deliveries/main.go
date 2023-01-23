package deliveries

import (
	"encoding/json"
	"fmt"
	"log"
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

func GetAllDeliveryDetails(page int) {
	db, err := utils.GetDBContext()
	if err != nil {
		log.Fatal(err.Error())
	}

	offset := 5 * (page - 1)

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
			LIMIT 5
			OFFSET ?
		`, offset).
		Find(&orderDetails)

	PrintJSON(orderDetails)
}

func GetDeliveryDetailsByDate(startDate time.Time, endDate time.Time, page int) {
	db, err := utils.GetDBContext()
	if err != nil {
		log.Fatal(err.Error())
	}

	offset := 5 * (page - 1)

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
			orders.created_at between ? and ?
		LIMIT 5
		OFFSET ?`, startDate, endDate, offset).
		Find(&orderDetails)

	PrintJSON(orderDetails)
}

func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}
