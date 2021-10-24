package library

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zahradnik.xyz/ksp-entertainment/database"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", GetLibraryItems)
}

func GetLibraryItems(c *gin.Context) {
	pageString, exists := c.GetQuery("page")
	page := 1

	if exists {
		p, err := strconv.Atoi(pageString)
		page = p
		if err != nil {
			c.Status(404)
			return
		}
	}

	var items []database.LibraryItem
	database.DB.Limit(100).Offset((page - 1) * 100).Find(&items)
	c.JSON(200, gin.H{"items": items})
}
