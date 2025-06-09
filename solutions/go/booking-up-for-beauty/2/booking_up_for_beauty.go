package booking

import (
    "time"
)

func Schedule(date string) time.Time {
	v, err := time.Parse("1/2/2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
    return v
}

func HasPassed(date string) bool {
    v, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return v.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
    v, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return v.Hour() >= 12 && v.Hour() < 18
}

func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
