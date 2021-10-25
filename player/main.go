package player

import (
	"log"
	"net"
	"os/exec"
	"time"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/parsers"
)

var MpvRunning = false
var MpvPaused = false
var QueueStopped = false

func RunPlayerWorker() {
	for {
		var queueItem database.QueueItem
		res := database.DB.Where("played_at IS NULL").Order("id").Preload("LibraryItem").Take(&queueItem)
		if res.RowsAffected == 0 || QueueStopped {
			// Empty queue
			time.Sleep(5 * time.Second)
			continue
		}

		queueItem.PlayedAt.Time = time.Now()
		queueItem.PlayedAt.Valid = true
		database.DB.Save(&queueItem)

		queueItem.LibraryItem.PlayCount++
		database.DB.Save(&queueItem.LibraryItem)

		err := StartMpv(parsers.ParseURL(queueItem.LibraryItem.URL))
		if err != nil {
			log.Printf("Error from MPV process: %v\n", err)
		}
	}
}

func StartMpv(url string) error {
	log.Printf("Starting MPV for URL %v\n", url)
	MpvRunning = true
	MpvPaused = false
	cmd := exec.Command("/usr/bin/mpv", url, "--input-ipc-server=/tmp/mpv.sock", "--no-video")
	err := cmd.Run()
	MpvRunning = false
	return err
}

func SendMpvCommand(command string) error {
	c, err := net.Dial("unix", "/tmp/mpv.sock")
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Write([]byte(command + "\n"))
	return err
}

func StopMpv() error {
	if !MpvRunning {
		return nil
	}

	return SendMpvCommand("quit")
}

func PauseMpv() error {
	if !MpvRunning {
		return nil
	}

	MpvPaused = !MpvPaused
	return SendMpvCommand("cycle pause")
}
