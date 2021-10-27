package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zahradnik.xyz/ksp-entertainment/telegram"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", RedirectHome)
	router.GET("/telegram/", Telegram)
	router.GET("/queue/", QueueView)
	router.GET("/library/", LibraryIndex)
}

func RedirectHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/queue/")
}

func Telegram(c *gin.Context) {
	c.HTML(200, "telegram.gohtml", gin.H{
		"password": telegram.CurrentPassword,
	})
}
