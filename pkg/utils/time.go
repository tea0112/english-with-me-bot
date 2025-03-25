package utils

import (
	"fmt"
	"time"
)

func FormatToVietnamSheetTime(t time.Time) (day string, timeStr string) {
	// Load the Vietnam time zone.
	vietnamLocation, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Get the current time in the Vietnam time zone.
	t = t.In(vietnamLocation)

	day = t.Format("02/01/2006")   // Format as "DD/MM/YYYY"
	timeStr = t.Format("15:04:05") // Format as "hh:mm:ss"
	return day, timeStr
}
