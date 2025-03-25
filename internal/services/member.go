package services

import (
	"context"
	"log"

	"github.com/cesc1802/english-with-me-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MemberSvcImpl struct {
	sheetRepo SheetRepo
}

func NewMemberSvcImpl(
	sheetRepo SheetRepo,
) *MemberSvcImpl {
	return &MemberSvcImpl{
		sheetRepo: sheetRepo,
	}
}

func (s *MemberSvcImpl) SaveMember(ctx context.Context, incomingUpdate tgbotapi.Update) error {
	for _, member := range incomingUpdate.Message.NewChatMembers {
		// When new member joins, call function to save them
		err := s.sheetRepo.SaveNewMember(ctx, models.GroupMemberInfo{
			GroupUserId: member.ID,
			Username:    member.UserName,
			Fullname:    member.FirstName + " " + member.LastName,
		})

		if err != nil {
			log.Printf("[ERROR] Save new member error: %v", err)
			return err
		}
	}

	return nil
}
