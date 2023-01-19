package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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
	importOrders()

}

func importOrders() {
	const filePath = "seed_data/Test task - Postgres - orders.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read orders csv file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	// fmt.Printf("%v", records)

	for index, row := range records {
		// skip column names
		if index > 1 {
			// fmt.Println(row)
			fmt.Print(row[1], "\n")
		}
	}

	fmt.Printf("\n")
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
