package main

import (
	"github.com/spf13/viper"
	"log"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/player"
	"zahradnik.xyz/ksp-entertainment/web"
)

func main() {
	log.Println("Starting entertainment...")

	viper.SetDefault("app.port", 8001)
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

	web.RunWebServer(viper.GetInt("app.port"))
}
