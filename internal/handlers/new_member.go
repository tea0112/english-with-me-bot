package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *IncomingUpdateHandler) handleNewChatMembers(ctx context.Context, incomingUpdate tgbotapi.Update) {
	err := h.memberSvc.SaveMember(ctx, incomingUpdate)
	if err != nil {
		log.Printf("[ERROR] Topic service error: %v", err)
	}
}
