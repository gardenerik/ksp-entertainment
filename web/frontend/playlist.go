package frontend

import (
	"github.com/gin-gonic/gin"
	"zahradnik.xyz/ksp-entertainment/database"
)

func ListPlaylists(c *gin.Context)  {
	var playlists []database.Playlist
	database.DB.Order("name").Find(&playlists)

	c.HTML(200, "playlists.gohtml", gin.H{
		"playlists": playlists,
	})
}

func GetPlaylist(c *gin.Context)  {
	var playlist database.Playlist
	database.DB.Preload("LibraryItems").Find(&playlist, c.Param("id"))

	c.HTML(200, "playlist_detail.gohtml", gin.H{
		"playlist": playlist,
	})
}