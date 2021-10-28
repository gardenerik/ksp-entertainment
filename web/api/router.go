package api

import (
	"github.com/gin-gonic/gin"
	"zahradnik.xyz/ksp-entertainment/web/api/library"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/queue/", GetQueue)
	router.POST("/queue/", AddToQueue)
	router.DELETE("/queue/", ClearQueue)
	router.DELETE("/queue/:id/", DeleteFromQueue)

	router.GET("/playback/now/", NowPlaying)
	router.POST("/playback/skip/", SkipPlayback)
	router.POST("/playback/resume/", ResumePlayback)
	router.POST("/playback/stop/", StopPlayback)
	router.POST("/playback/pause/", PausePlayback)

	router.GET("/playlists/", ListPlaylists)
	router.POST("/playlists/", AddPlaylist)
	router.POST("/playlists/import/", ImportPlaylist)
	router.GET("/playlists/:id/", GetPlaylist)
	router.POST("/playlists/:id/", AddPlaylistItem)
	router.POST("/playlists/:id/play/", PlayPlaylist)

	library.RegisterRoutes(router.Group("/library"))
}
