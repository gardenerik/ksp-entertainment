package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
	"zahradnik.xyz/ksp-entertainment/database"
)

var bot *tgbotapi.BotAPI

func sendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("[Telegram] Error while sending message: %v\n", err)
	}
}

func sendDeniedMessage(chatId int64) {
	sendMessage(chatId, "🛑 STOP! I don't know you.\n\nSee instructions under Telegram Pairing on Entertainment.")
}

func StartTelegramBot() {
	token := viper.GetString("app.telegram_token")
	if token == "" {
		return
	}

	go RunPasswordLoop()

	b, err := tgbotapi.NewBotAPI(token)
	bot = b
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		text := update.Message.Text
		if strings.HasPrefix(text, "/start") {
			HandleStart(update)
			continue
		}

		var user database.TelegramUser
		res := database.DB.Where("id = ?", update.Message.From.ID).Where("verified_at >= ?", time.Now().Add(-time.Hour*24*14)).Find(&user)
		if res.RowsAffected == 0 {
			sendDeniedMessage(update.Message.Chat.ID)
			continue
		}

		if strings.HasPrefix(text, "http") {
			lib := database.GetOrAddLibraryItem(strings.TrimSpace(text), "")
			database.AddToQueue(lib)
			log.Printf("[Telegram] %v added %v (%v) to queue.\n", update.Message.From.UserName, lib.Name, lib.URL)
			sendMessage(update.Message.Chat.ID, "🎵 Added to queue.")
			continue
		}

		sendMessage(update.Message.Chat.ID, "I did not understand that.")
	}
}
