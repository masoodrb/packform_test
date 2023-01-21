package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"masoodahm/packform/models"
	"masoodahm/packform/utils"
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
	importCompanies(dbContext, "seed_data/Test task - Postgres - customer_companies.csv")
	importCustomers(dbContext, "seed_data/Test task - Postgres - customers.csv")
	importOrders(dbContext, "seed_data/Test task - Postgres - orders.csv")
	importOrderItems(dbContext, "seed_data/Test task - Postgres - order_items.csv")
	importDeliveries(dbContext, "seed_data/Test task - Postgres - deliveries.csv")

}

func importCompanies(db *gorm.DB, csvFilePath string) {
	file := openFile(csvFilePath)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	for index, row := range records {
		// skip first row as it is just column names
		if index >= 1 {
			company := models.Company{}

			company_id, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}
			company.Id = company_id
			company.CompanyName = row[1]
			result := db.Create(&company)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func importCustomers(db *gorm.DB, csvFilePath string) {
	file := openFile(csvFilePath)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	for index, row := range records {
		// skip first row as it is just column names
		if index >= 1 {
			customer := models.Customer{}
			customer.UserId = row[0]
			customer.Login = row[1]
			password, err := utils.HashPassword(row[2])
			if err != nil {
				log.Fatal(err.Error())
			}
			customer.Password = password
			customer.Name = row[3]
			customer.CreditCards = row[5]
			companyId, err := strconv.ParseInt(row[4], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}
			customer.CompanyId = companyId
			result := db.Create(&customer)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func importOrders(db *gorm.DB, csvFilePath string) {
	file := openFile(csvFilePath)
	defer file.Close()

	csvReader := csv.NewReader(file)
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
				createdAt = time.Now()
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

func importDeliveries(db *gorm.DB, csvFilePath string) {
	file := openFile(csvFilePath)
	defer file.Close()

	csvReader := csv.NewReader(file)
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
				delivered_quantity = 0
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

func importOrderItems(db *gorm.DB, csvFilePath string) {
	f := openFile(csvFilePath)
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
				unit_price = 0
			}

			quantity, err := strconv.ParseInt(row[3], 10, 64)
			if err != nil {
				quantity = 0
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
		&models.Company{},
		&models.Customer{},
		&models.Delivery{},
		&models.Order{},
		&models.OrderItem{},
	)
}

func getDBContext(db_host string, db_name string, db_pass string, db_port string, db_user string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Australia/Melbourne",
		db_host, db_user, db_pass, db_name, db_port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
