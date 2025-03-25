package repositories

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cesc1802/english-with-me-bot/internal/models"
	"github.com/cesc1802/english-with-me-bot/pkg/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/sheets/v4"
)

type SheetRepoImpl struct {
	sheetSvc               *sheets.Service
	spreadsheetFileId      string
	announcementSheetName  string
	submitSheetName        string
	contactMemberSheetName string
}

func NewSheetRepoImpl(
	sheetSvc *sheets.Service,
	spreadsheetFileId string,
	announcementSheetName string,
	submitSheetName string,
	contactMemberSheetName string,
) *SheetRepoImpl {
	return &SheetRepoImpl{
		sheetSvc:               sheetSvc,
		spreadsheetFileId:      spreadsheetFileId,
		announcementSheetName:  announcementSheetName,
		submitSheetName:        submitSheetName,
		contactMemberSheetName: contactMemberSheetName,
	}
}

func (r *SheetRepoImpl) SaveMessageToAnnoucementSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error {
	submitDay, submitTime := utils.FormatToVietnamSheetTime(time.Now())
	username := tgBotAPIMessage.From.UserName
	fullname := fmt.Sprintf("%s %s", tgBotAPIMessage.From.FirstName, tgBotAPIMessage.From.LastName)
	text := tgBotAPIMessage.Text

	annoucement := models.AnnoucementSheet{
		Day:                submitDay,
		Time:               submitTime,
		Username:           username,
		Fullname:           fullname,
		AnnoucementContent: text,
	}

	values := [][]interface{}{
		annoucement.ToSheetValue(),
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := r.sheetSvc.Spreadsheets.Values.Append(
		r.spreadsheetFileId,
		fmt.Sprintf("%s!A1:E1", r.announcementSheetName),
		valueRange,
	).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

	if err != nil {
		return err
	}

	return nil
}

func (r *SheetRepoImpl) SaveMessageToSubmitSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error {
	submitDay, submitTime := utils.FormatToVietnamSheetTime(time.Now())
	username := tgBotAPIMessage.From.UserName
	fullname := fmt.Sprintf("%s %s", tgBotAPIMessage.From.FirstName, tgBotAPIMessage.From.LastName)
	text := tgBotAPIMessage.Text

	submitSheet := models.SubmitSheet{
		Day:           submitDay,
		Time:          submitTime,
		Username:      username,
		Fullname:      fullname,
		SubmitContent: text,
	}

	values := [][]interface{}{
		submitSheet.ToSheetValue(),
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := r.sheetSvc.Spreadsheets.Values.Append(
		r.spreadsheetFileId,
		fmt.Sprintf("%s!A1:E1", r.submitSheetName),
		valueRange,
	).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

	if err != nil {
		return err
	}

	return nil
}

func (r *SheetRepoImpl) SaveNewMember(ctx context.Context, memberInfo models.GroupMemberInfo) error {
	// Range to check for existing entries
	readRange := "contact_member!B2:Z2"

	// Check if user already exists by username
	resp, err := r.sheetSvc.Spreadsheets.Values.Get(r.spreadsheetFileId, readRange).Do()
	if err != nil {
		log.Printf("Unable to retrieve data: %v", err)
		return err
	}

	userExists := false
	for _, row := range resp.Values {
		if val, ok := row[0].(string); ok && val != "" {
			i, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("Unable to convert GroupUserId to int: %v", err)
				return err
			}

			if int64(i) == memberInfo.GroupUserId {
				userExists = true
				break
			}
		} else {
			// Handle empty cell or non-string value - skip this row
			continue
		}
	}

	// If user does not exist, add them to the sheet
	if !userExists {
		// Prepare the values to be inserted
		var values [][]interface{}
		values = append(values, []interface{}{memberInfo.GroupUserId, memberInfo.Username, memberInfo.Fullname})

		// Append values to the sheet
		valueRange := &sheets.ValueRange{
			Values: values,
		}

		_, err := r.sheetSvc.Spreadsheets.Values.Append(
			r.spreadsheetFileId,
			fmt.Sprintf("%s!A1:E1", r.contactMemberSheetName),
			valueRange,
		).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

		if err != nil {
			log.Printf("Unable to write data: %v", err)
			return err
		}

		log.Printf("Added new member: %s (%s)", memberInfo.Username, memberInfo.Fullname)
		return nil
	}

	log.Printf("Member already exists: %s (%s)", memberInfo.Username, memberInfo.Fullname)
	return nil
}
