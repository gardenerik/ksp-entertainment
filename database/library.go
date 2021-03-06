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
	Name         string `json:"name"`
	LibraryItems []LibraryItem `gorm:"many2many:playlist_items;" json:"library_items"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetOrAddLibraryItem(url, name string) LibraryItem {
	var lib LibraryItem
	res := DB.Where("url = ?", url).Limit(1).Find(&lib)
	if res.RowsAffected == 0 {
		if name != "" {
			lib.Name = name
		} else {
			lib.Name = parsers.GetName(url)
		}
		lib.URL = url
		DB.Save(&lib)
	}

	return lib
}
