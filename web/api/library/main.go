package library

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"zahradnik.xyz/ksp-entertainment/database"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", GetLibraryItems)
	router.POST("/", AddLibraryItem)
	router.POST("/:id/enqueue/", EnqueueItem)
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

func EnqueueItem(c *gin.Context) {
	id := c.Param("id")
	var item database.LibraryItem
	res := database.DB.First(&item, id)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return
	}

	database.AddToQueue(item)
	c.Status(200)
}

func AddLibraryItem(c *gin.Context) {
	var data struct {
		URL  string `form:"url"`
		Name string `form:"name"`
	}

	err := c.ShouldBind(&data)
	if err != nil {
		c.Status(400)
		return
	}

	data.URL = strings.TrimSpace(data.URL)
	data.Name = strings.TrimSpace(data.Name)

	if data.URL == "" {
		c.Status(400)
		return
	}

	item := database.GetOrAddLibraryItem(data.URL, data.Name)
	c.JSON(200, item)
}
