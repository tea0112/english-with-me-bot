package sheet

import (
	"context"
	"log"

	"github.com/cesc1802/english-with-me-bot/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSheetConn(config *config.AppConfig) sheets.Service {
	credsPath := "sheet.json" // Path to your service account JSON key file
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		log.Fatalf("Unable to create Sheets client: %v", err)
	}

	return *srv

	//readRange := "announcement!A1:E1"
	//srv.Spreadsheets.Values.Get(os.Getenv("SHEET_ID"), readRange)
	//resp, err := srv.Spreadsheets.Values.Get(os.Getenv("SHEET_ID"), readRange).Do()
	//if err != nil {
	//	log.Fatalf("Unable to read data from sheet: %v", err)
	//}
	//
	//if len(resp.Values) == 0 {
	//	fmt.Println("No data found.")
	//} else {
	//	for _, row := range resp.Values {
	//		fmt.Println(row)
	//	}
	//}
}
