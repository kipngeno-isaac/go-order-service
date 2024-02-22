package initializers

import (
	"log"
	"os"

	"github.com/kipngeno-isaac/go-order-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to databse")
	}
	DB.AutoMigrate(&models.Customer{}, &models.Order{})
}
