package api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/player"
)

func NowPlaying(c *gin.Context) {
	var now_playing database.QueueItem
	res := database.DB.Where("played_at IS NOT NULL").Order("played_at desc").Preload("LibraryItem").Take(&now_playing)
	if res.RowsAffected == 0 {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, gin.H{
		"name":     now_playing.LibraryItem.Name,
		"url":      now_playing.LibraryItem.URL,
		"start_at": now_playing.PlayedAt,
		"added_by": now_playing.AddedBy,
	})
}

func StopPlayer(c *gin.Context) {
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

func PausePlayer(c *gin.Context) {
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

func AddToQueue(c *gin.Context) {
	url, present := c.GetPostForm("url")
	if !present || strings.TrimSpace(url) == "" {
		c.JSON(400, gin.H{"error": "Bad url in request data."})
		return
	}

	adder, present := c.GetPostForm("adder")
	if !present || strings.TrimSpace(url) == "" {
		c.JSON(400, gin.H{"error": "Bad adder in request data."})
		return
	}

	lib := database.GetOrAddLibraryItem(url)
	database.AddToQueue(lib, "api:"+adder)
	c.JSON(200, lib)
}
