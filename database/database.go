package database

import (
	"fmt"
	"log"
	"os"
	"perpustakaan/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// mewakili dari instance dari db yang kita pakai untuk opeasi berbagai query
var DB *gorm.DB

// connect ke database
func InitDatabase() {
	var (
		DB_USERNAME string = os.Getenv("DB_USERNAME")
		DB_PASSOWRD string = os.Getenv("DB_PASSWORD")
		DB_NAME     string = os.Getenv("DB_NAME")
		DB_HOST     string = os.Getenv("DB_HOST")
		DB_PORT     string = os.Getenv("DB_PORT")
	)
	var err error
	// dsn (data source name) alamat yang akan mengarah ke database kita
	//?charset=utf8mb4&parseTime=true&loc=Local
	var dsn string = fmt.Sprintf("%s:%s@%s:%s/%s",
		DB_USERNAME,
		DB_PASSOWRD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	fmt.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")
}

// perform migration

func Migrate() {
	err := DB.AutoMigrate(&models.Book{}, &models.Member{}, &models.Admin{}, &models.Transaction{})

	if err != nil {
		log.Fatalf("failed to perform books database migration: %s\n", err)
	}
}
