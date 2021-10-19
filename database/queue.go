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
	PlayedAt  null.Time `gorm:"index"`
	//Order uint `gorm:"index"`
	AddedBy       string
	LibraryItemID uint
	LibraryItem   LibraryItem
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func AddToQueue(lib LibraryItem, adder string) {
	DB.Save(&QueueItem{
		PlayedAt:      null.Time{},
		AddedBy:       adder,
		LibraryItemID: lib.ID,
	})
}
