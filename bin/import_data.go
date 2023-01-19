package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"masoodahm/packform/models"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	db_user := goDotEnvVariable("POSTGRES_USER")
	db_pass := goDotEnvVariable("POSTGRES_PASSWORD")
	db_name := goDotEnvVariable("POSTGRES_DB")
	db_port := goDotEnvVariable("POSTGRES_PORT")
	db_host := goDotEnvVariable("DB_HOST")

	dbContext, err := getDBContext(db_host, db_name, db_pass, db_port, db_user)
	if err != nil {
		log.Fatalln(err)
	}

	createTables(dbContext)
	importOrders(dbContext, "seed_data/Test task - Postgres - orders.csv")
	importOrderItems(dbContext, "seed_data/Test task - Postgres - order_items.csv")
	importDeliveries(dbContext, "Test task - Postgres - deliveries.csv")
}

func importOrders(db *gorm.DB, filePath string) {
	f := openFile(filePath)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	for index, row := range records {
		// skip first row as it is just column names
		if index >= 1 {
			id, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}
			createdAt, err := time.Parse(time.RFC3339, row[1])
			if err != nil {
				log.Fatal(err.Error())
			}
			order := models.Order{}
			order.CreatedAt = createdAt
			order.Id = id
			order.OrderName = row[2]
			order.CustomerId = row[3]
			result := db.Create(&order)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func importDeliveries(db *gorm.DB, filePath string) {
	f := openFile(filePath)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	for index, row := range records {
		// skip first row as it is just column names
		if index >= 1 {
			id, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			order_item_id, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			delivered_quantity, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			delivery := models.Delivery{}
			delivery.Id = id
			delivery.OrderItemId = order_item_id
			delivery.DeliveredQuantity = delivered_quantity
			result := db.Create(&delivery)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func openFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read orders csv file")
	}

	return f
}

func importOrderItems(db *gorm.DB, filePath string) {
	f := openFile(filePath)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	for index, row := range records {
		// skip first row as it is just column names
		if index >= 1 {
			id, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}
			order_id, err := strconv.ParseInt(row[1], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			unit_price, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			quantity, err := strconv.ParseInt(row[3], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			orderItem := models.OrderItem{}
			orderItem.Id = id
			orderItem.OrderId = order_id
			orderItem.PricePerUnit = unit_price
			orderItem.Quantity = quantity
			orderItem.Product = row[4]
			result := db.Create(&orderItem)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func createTables(db *gorm.DB) {

	db.AutoMigrate(
		&models.CustomerComapnies{},
		&models.Customers{},
		&models.Delivery{},
		&models.Order{},
		&models.OrderItem{},
	)
}

func getDBContext(db_host string, db_name string, db_pass string, db_port string, db_user string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Australia/Melbourne", db_host, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
