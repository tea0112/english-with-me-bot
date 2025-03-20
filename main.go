package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/cesc1802/english-with-me-bot/config"
	"github.com/cesc1802/english-with-me-bot/statics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func prepareSpreadsheet(srv *sheets.Service, spreadsheetId string) {
	// Check if headers exist, if not add them
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, "announcement!A1:E1").Do()
	if err != nil {
		log.Fatalf("Unable to read sheet headers: %v", err)
	}

	if len(resp.Values) == 0 {
		headers := [][]interface{}{
			{"Date", "Time", "Username", "Name Of User", "Message Text"},
		}

		valueRange := &sheets.ValueRange{
			Values: headers,
		}

		_, err := srv.Spreadsheets.Values.Update(
			spreadsheetId,
			"announcement!A1:E1",
			valueRange,
		).ValueInputOption("RAW").Do()

		if err != nil {
			log.Fatalf("Unable to write headers: %v", err)
		}
	}
}

func saveMessageToSheet(srv *sheets.Service, spreadsheetId string, message *tgbotapi.Message, channelSheetName string) {
	timestamp := time.Now().Format(time.RFC3339)
	username := message.From.UserName
	if username == "" {
		username = fmt.Sprintf("%s %s", message.From.FirstName, message.From.LastName)
	}

	values := [][]interface{}{
		{
			timestamp,
			timestamp,
			username,
			message.From.ID,
			message.Text,
		},
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := srv.Spreadsheets.Values.Append(
		spreadsheetId,
		fmt.Sprintf("%s!A1:E1", channelSheetName),
		valueRange,
	).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

	if err != nil {
		log.Printf("Error saving to sheet: %v", err)
	} else {
		log.Printf("Successfully saved message from %s to Google Sheets", username)
	}
}

func selectEnvFile() (string, string) {
	if len(os.Args) < 2 {
		log.Fatal("Please provide an environment argument: 'dev' or 'prod'")
	}

	serviceEnv := os.Args[1]
	var envFile string
	switch serviceEnv {
	case statics.SERVICE_ENV_DEV:
		envFile = ".env.dev"
	case statics.SERVICE_ENV_PROD:
		envFile = ".env.prod"
	default:
		log.Fatal("Invalid environment. Use 'dev' or 'prod'")
	}

	return envFile, serviceEnv
}

func main() {
	// Initialize environment variables
	envFile, serviceEnv := selectEnvFile()
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
	fmt.Printf("Loaded environment from %s\n", envFile)

	var appCfg config.AppConfig
	if err := env.Parse(&appCfg); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(appCfg.BotToken)
	if err != nil {
		log.Fatalf("Error starting Telegram bot: %v", err)
	}

	if serviceEnv == statics.SERVICE_ENV_DEV {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// google sheet
	credBytes, err := base64.StdEncoding.DecodeString(appCfg.GoogleSheetCredsBase64)
	if err != nil {
		log.Fatalf("Failed to decode credentials: %v", err)
	}
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Failed to create JWT config: %v", err)
	}

	client := config.Client(context.Background())
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Failed to create Sheets service: %v", err)
	}

	// Ensure spreadsheet has the right headers
	prepareSpreadsheet(srv, appCfg.SpreadsheetId)

	//context
	//botCtx := botcontext.NewBotContext(bot, &appCfg)
	// Set up the update listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// Process incoming messages
	for update := range updates {
		if update.Message != nil && strings.Contains(update.Message.Text, "#topic") {
			saveMessageToSheet(srv, appCfg.SpreadsheetId, update.Message, "announcement")

			// Respond to the message
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Saved Topic!")
			reply.ReplyToMessageID = appCfg.AnnouncementsTopicId
			bot.Send(reply)
		}

		if update.Message != nil && strings.Contains(update.Message.Text, "#submit") {
			saveMessageToSheet(srv, appCfg.SpreadsheetId, update.Message, "submit")

			// Respond to the message
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Submitted!")
			reply.ReplyToMessageID = appCfg.StudentPresentationsTopicId
			bot.Send(reply)
		}

	}
}
