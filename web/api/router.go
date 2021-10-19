package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/now_playing", NowPlaying)
	router.POST("/stop", StopPlayer)
	router.POST("/pause", PausePlayer)
	router.POST("/add_to_queue", AddToQueue)
}
