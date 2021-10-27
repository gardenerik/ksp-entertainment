package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open(viper.GetString("app.database")), &gorm.Config{
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
	DB.AutoMigrate(&TelegramUser{})

	log.Println("Database ready!")
}
