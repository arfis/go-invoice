package database

import (
	"fmt"
	"github.com/arfis/go-invoice/authorization/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	connection *gorm.DB
	once       sync.Once
)

type Database struct {
	Val int
}

func (database *Database) TrySelect() {
	connection.AutoMigrate(&model.User{})

	connection.Create(&model.User{Login: "D42", Password: 100})

	// Read
	var user model.User
	result := connection.First(&user, 1) // retrieves the first Product
	if result.Error != nil {
		log.Fatalf("Error when getting product: %v", result.Error)
	}
	fmt.Printf("Product found: %s, %d\n", user.Login, user.Password)
}

func (database Database) Test() {

}

func (database Database) Test2() Database {
	database.Val = 10
	return database
}

func (database *Database) Test3() {
	database.Val = 12
}

func (database *Database) CreateConnection() *gorm.DB {
	once.Do(func() {
		host := os.Getenv("DB_HOST")         // Get the host from environment variables
		port := os.Getenv("DB_PORT")         // Get the port from environment variables
		user := os.Getenv("DB_USER")         // Get the user from environment variables
		password := os.Getenv("DB_PASSWORD") // Get the password from environment variables
		dbname := os.Getenv("DB_NAME")       // Get the database name from environment variables
		sslmode := "disable"
		timeZone := "Europe/Bratislava" // Changed from "Asia/Shanghai" to "Europe/Bratislava"

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			host, user, password, dbname, port, sslmode, timeZone)

		// Connect to the database using GORM
		connectedDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		} else {
			fmt.Println("Connected to the database successfully.")
		}

		connection = connectedDb
	})

	return connection
}
