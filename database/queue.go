package database

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type QueueItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PlayedAt  null.Time `gorm:"index" json:"played_at"`
	//Order uint `gorm:"index"`
	AddedBy       string         `json:"added_by"`
	LibraryItemID uint           `json:"library_item_id"`
	LibraryItem   LibraryItem    `json:"library_item,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func AddToQueue(lib LibraryItem, adder string) {
	DB.Save(&QueueItem{
		PlayedAt:      null.Time{},
		AddedBy:       adder,
		LibraryItemID: lib.ID,
	})
}

func GetQueue() {

}
