package api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/player"
)

func NowPlaying(c *gin.Context) {
	var nowPlaying database.QueueItem
	res := database.DB.Where("played_at IS NOT NULL").Order("played_at desc").Preload("LibraryItem").Take(&nowPlaying)
	if res.RowsAffected == 0 {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, gin.H{
		"current":       nowPlaying,
		"queue_stopped": player.QueueStopped,
		"playing":       player.MpvRunning,
		"paused":        player.MpvPaused,
	})
}

func StopPlayback(c *gin.Context) {
	if !player.MpvRunning {
		c.Status(200)
		return
	}

	player.QueueStopped = true
	err := player.StopMpv()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.Status(200)
}

func SkipPlayback(c *gin.Context) {
	if !player.MpvRunning {
		c.Status(200)
		return
	}

	err := player.StopMpv()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.Status(200)
}

func PausePlayback(c *gin.Context) {
	if !player.MpvRunning {
		c.Status(200)
		return
	}

	err := player.PauseMpv()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.Status(200)
}

func ResumePlayback(c *gin.Context) {
	player.QueueStopped = false
	c.Status(200)
}

func AddToQueue(c *gin.Context) {
	url, present := c.GetPostForm("url")
	if !present || strings.TrimSpace(url) == "" {
		c.JSON(400, gin.H{"error": "Bad url in request data."})
		return
	}

	lib := database.GetOrAddLibraryItem(url, "")
	database.AddToQueue(lib)
	c.JSON(200, lib)
}

func GetQueue(c *gin.Context) {
	var items []database.QueueItem
	database.DB.Where("played_at IS NULL").Order("id").Preload("LibraryItem").Find(&items)
	c.JSON(200, gin.H{"queue": items})
}

func ClearQueue(c *gin.Context) {
	database.DB.Where("played_at IS NULL").Delete(&database.QueueItem{})
	c.Status(200)
}

func DeleteFromQueue(c *gin.Context) {
	itemId := c.Param("id")
	database.DB.Where("played_at IS NULL").Where("id = ?", itemId).Delete(&database.QueueItem{})
	c.Status(200)
}
