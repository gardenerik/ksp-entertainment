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

	library.RegisterRoutes(router.Group("/library"))
}
