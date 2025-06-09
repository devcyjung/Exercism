package booking

import (
    "time"
)

func Schedule(date string) time.Time {
	v, _ := time.Parse("1/2/2006 15:04:05", date)
    return v
}

func HasPassed(date string) bool {
    v, _ := time.Parse("January 2, 2006 15:04:05", date)
	return v.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
    v, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return v.Hour() >= 12 && v.Hour() < 18
}

func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
