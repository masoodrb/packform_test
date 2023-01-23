package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	db_user := GoDotEnvVariable("POSTGRES_USER")
	db_pass := GoDotEnvVariable("POSTGRES_PASSWORD")
	db_name := GoDotEnvVariable("POSTGRES_DB")
	db_port := GoDotEnvVariable("POSTGRES_PORT")
	db_host := GoDotEnvVariable("DB_HOST")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Australia/Melbourne",
		db_host, db_user, db_pass, db_name, db_port,
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error creating database connection")
	}
}

func GetDBContext() *gorm.DB {
	return db
}
