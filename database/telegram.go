package database

import "time"

type TelegramUser struct {
	ID         int64     `gorm:"primarykey" json:"id"`
	VerifiedAt time.Time `json:"verified_at"`
	Name       string    `json:"name"`
}
