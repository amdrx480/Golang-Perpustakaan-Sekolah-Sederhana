package database

import (
	"fmt"
	"log"
	"perpustakaan/models"
	"perpustakaan/utils"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// mewakili dari instance dari db yang kita pakai untuk opeasi berbagai query
var DB *gorm.DB

var (
	DB_USERNAME string = utils.GetConfig("DB_USERNAME")
	DB_PASSOWRD string = utils.GetConfig("DB_PASSWORD")
	DB_NAME     string = utils.GetConfig("DB_NAME")
	DB_HOST     string = utils.GetConfig("DB_HOST")
	DB_PORT     string = utils.GetConfig("DB_PORT")
)

// connect ke database
func InitDatabase() {
	var err error
	// dsn (data source name) alamat yang akan mengarah ke database kita
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		DB_USERNAME,
		DB_PASSOWRD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
