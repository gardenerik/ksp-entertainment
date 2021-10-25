package frontend

import (
	"github.com/gin-gonic/gin"
)

func QueueView(c *gin.Context) {
	c.HTML(200, "queue.gohtml", nil)
}
