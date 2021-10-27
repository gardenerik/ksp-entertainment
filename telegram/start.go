package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
	"time"
	"zahradnik.xyz/ksp-entertainment/database"
)

func HandleStart(update tgbotapi.Update) {
	parts := strings.Split(update.Message.Text, " ")
	if len(parts) != 2 {
		sendDeniedMessage(update.Message.Chat.ID)
		return
	}

	if parts[1] == CurrentPassword {
		user := database.TelegramUser{
			ID:         int64(update.Message.From.ID),
			VerifiedAt: time.Now(),
			Name:       update.Message.From.UserName,
		}
		database.DB.Save(&user)

		sendMessage(update.Message.Chat.ID, "✅ Successfully verified.")
		log.Printf("[Telegram] Verified %v.\n", update.Message.From.UserName)
	} else {
		sendMessage(update.Message.Chat.ID, "⛔️ Wrong password.")
	}
}
