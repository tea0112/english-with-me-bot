package main

import (
	"log"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/cesc1802/english-with-me-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var appCfg config.AppConfig
	if err := env.Parse(&appCfg); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(appCfg.BotToken)
	if err != nil {
		log.Fatalf("Error starting Telegram bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//context
	//botCtx := botcontext.NewBotContext(bot, &appCfg)
	// Set up the update listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// Process incoming messages
	for update := range updates {
		if update.Message != nil && strings.Contains(update.Message.Text, "@qanda20250315bot") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "nice to meet you")
			reply.ReplyToMessageID = appCfg.AnnouncementsTopicId
			bot.Send(reply)
		}

		if update.Message != nil && strings.Contains(update.Message.Text, "#submit") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "cam on ban da nop bai tap")
			reply.ReplyToMessageID = appCfg.StudentPresentationsTopicId
			bot.Send(reply)
		}

	}
}
