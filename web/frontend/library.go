package frontend

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zahradnik.xyz/ksp-entertainment/database"
)

func LibraryIndex(c *gin.Context) {
	var items []database.LibraryItem
	pageStr, pagePresent := c.GetQuery("page")
	page := 1
	if pagePresent {
		p, err := strconv.Atoi(pageStr)
		page = p
		if err != nil {
			c.Status(404)
			return
		}
	}

	database.DB.Order("name").Limit(100).Offset((page - 1) * 100).Find(&items)

	var nextPage, prevPage int
	if len(items) == 100 {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	}

	c.HTML(200, "library.gohtml", gin.H{
		"items":    items,
		"nextPage": nextPage,
		"prevPage": prevPage,
	})
}
