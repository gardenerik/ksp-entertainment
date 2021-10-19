package database

import (
	"time"
	"zahradnik.xyz/ksp-entertainment/parsers"
)

type LibraryItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	URL       string    `gorm:"unique" json:"url"`
	PlayCount uint      `json:"play_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Playlist struct {
	ID           uint `gorm:"primarykey" json:"id"`
	Name         string
	LibraryItems []LibraryItem
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetOrAddLibraryItem(url string) LibraryItem {
	var lib LibraryItem
	res := DB.Where("url = ?", url).Limit(1).Find(&lib)
	if res.RowsAffected == 0 {
		lib.Name = parsers.GetName(url)
		lib.URL = url
		DB.Save(&lib)
	}

	return lib
}
