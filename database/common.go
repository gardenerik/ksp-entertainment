package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("entertainment.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				IgnoreRecordNotFoundError: true,
			},
		),
	})
	DB = db
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&LibraryItem{})
	DB.AutoMigrate(&QueueItem{})

	log.Println("Database ready!")
}
