package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", RedirectHome)
	router.GET("/queue/", QueueView)
}

func RedirectHome(c *gin.Context)  {
	c.Redirect(http.StatusFound, "/queue/")
}
