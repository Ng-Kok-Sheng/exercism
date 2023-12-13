package booking

import (
	"fmt"
	"time"
)

func ParseTimeWithLayout(date, layout string) time.Time {
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	return parsedTime
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	return ParseTimeWithLayout(date, layout)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedTime := ParseTimeWithLayout(date, layout)
	return time.Now().After(parsedTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime := ParseTimeWithLayout(date, layout)

	if parsedTime.Hour() >= 12 && parsedTime.Hour() <= 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime := ParseTimeWithLayout(date, "1/2/2006 15:04:05")
	formattedTime := appointmentTime.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedTime := ParseTimeWithLayout("2023-09-15 00:00:00", "2006-01-02 15:04:05")
	return parsedTime.UTC()

}
