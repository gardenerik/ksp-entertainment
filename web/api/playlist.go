package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"zahradnik.xyz/ksp-entertainment/database"
	"zahradnik.xyz/ksp-entertainment/parsers"
)

func ListPlaylists(c *gin.Context) {
	playlists := []database.Playlist{}
	database.DB.Find(&playlists)

	c.JSON(200, gin.H{
		"playlists": playlists,
	})
}

func GetPlaylist(c *gin.Context)  {
	var playlist database.Playlist
	res := database.DB.Preload("LibraryItems").Find(&playlist, c.Param("id"))
	if res.RowsAffected == 0 {
		c.Status(404)
		return
	}

	c.JSON(200, playlist)
}

func AddPlaylist(c *gin.Context)  {
	name, present := c.GetPostForm("name")
	if !present || strings.TrimSpace(name) == "" {
		c.JSON(400, gin.H{"error": "Enter playlist name."})
		return
	}


	playlist := database.Playlist{
		Name: name,
	}
	database.DB.Save(&playlist)

	c.JSON(200, playlist)
}

func AddPlaylistItem(c *gin.Context)  {
	url, present := c.GetPostForm("url")
	if !present || strings.TrimSpace(url) == "" {
		c.JSON(400, gin.H{"error": "Enter item URL."})
		return
	}

	var playlist database.Playlist
	res := database.DB.Find(&playlist, c.Param("id"))
	if res.RowsAffected == 0 {
		c.Status(404)
		return
	}

	lib := database.GetOrAddLibraryItem(url, "")
	err := database.DB.Model(&playlist).Association("LibraryItems").Append(&lib)
	if err != nil {
		c.Status(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, lib)
}

func ImportPlaylist(c *gin.Context)  {
	url, present := c.GetPostForm("url")
	if !present || strings.TrimSpace(url) == "" {
		c.JSON(400, gin.H{"error": "Enter playlist URL."})
		return
	}

	name, present := c.GetPostForm("name")
	if !present || strings.TrimSpace(name) == "" {
		c.JSON(400, gin.H{"error": "Enter playlist name."})
		return
	}

	playlist := database.Playlist{
		Name: name,
	}
	database.DB.Save(&playlist)

	items := parsers.YoutubeDLParser{}.GetPlaylistList(url)
	for _, item := range items {
		lib := database.GetOrAddLibraryItem(item.URL(), item.Title)
		err := database.DB.Model(&playlist).Association("LibraryItems").Append(&lib)
		if err != nil {
			c.Status(500)
			fmt.Println(err)
		}
	}

	c.JSON(200, gin.H{
		"data": items,
	})
}

func PlayPlaylist(c *gin.Context)  {
	var playlist database.Playlist
	res := database.DB.Preload("LibraryItems").Find(&playlist, c.Param("id"))
	if res.RowsAffected == 0 {
		c.Status(404)
		return
	}

	for _, item := range playlist.LibraryItems {
		database.AddToQueue(item)
	}

	c.Status(200)
}