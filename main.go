package main

import (
	"log"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/player"
	"zahradnik.xyz/ksp-entertainment/web"
)

func main() {
	log.Println("Starting entertainment...")
	database.ConnectDatabase()
	go player.RunPlayerWorker()

	web.RunWebServer(8001)
}
