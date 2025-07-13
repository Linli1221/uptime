package database

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"uptime/backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_URI")

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	if dsn == "" {
		log.Println("DATABASE_URI is not set, falling back to sqlite.")
		DB, err = gorm.Open(sqlite.Open("test.db"), gormConfig)
	} else {
		// Parse the aiven-provided URI
		parsedURL, err := url.Parse(dsn)
		if err != nil {
			log.Fatal("Failed to parse DATABASE_URI!", err)
		}
		
		// Extract components
		password, _ := parsedURL.User.Password()
		host := parsedURL.Host
		path := strings.Trim(parsedURL.Path, "/")

		// Format for gorm
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			parsedURL.User.Username(),
			password,
			host,
			path,
		)

		DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	}

	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB.AutoMigrate(&models.Monitor{}, &models.Latency{}, &models.User{})
}