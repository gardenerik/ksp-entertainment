package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/player"
	"zahradnik.xyz/ksp-entertainment/telegram"
	"zahradnik.xyz/ksp-entertainment/web"
)

func main() {
	log.Println("Starting entertainment...")

	viper.SetDefault("app.port", 8001)
	viper.SetDefault("app.telegram_token", "")
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.database", "entertainment.db")
	viper.SetDefault("binaries.youtube_dl", "/usr/bin/youtube-dl")
	viper.SetDefault("binaries.mpv", "/usr/bin/mpv")

	viper.SetConfigFile("config.toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	database.ConnectDatabase()
	go player.RunPlayerWorker()

	if !viper.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	go telegram.StartTelegramBot()
	web.RunWebServer(viper.GetInt("app.port"))
}
