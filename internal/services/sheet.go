package services

import (
	"context"

	"github.com/cesc1802/english-with-me-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SheetRepo interface {
	SaveMessageToAnnoucementSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error
	SaveMessageToSubmitSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error
	SaveNewMember(ctx context.Context, memberInfo models.GroupMemberInfo) error
}
