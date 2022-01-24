//Package booking schedules an appointment for a beauty salon
package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	newDate, _ := time.Parse("1/2/2006 15:04:05", date)

	return time.Date(newDate.Year(), newDate.Month(), newDate.Day(), newDate.Hour(), newDate.Minute(), newDate.Second(), 0, time.UTC)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	newDate, _ := time.Parse("January 2, 2006 15:04:05", date)

	return time.Now().After(newDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	newDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return newDate.Hour() >= 12 && newDate.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	newDate, _ := time.Parse("1/2/2006 15:04:05", date)

	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", newDate.Weekday(), newDate.Month(), newDate.Day(), newDate.Year(), newDate.Hour(), newDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
