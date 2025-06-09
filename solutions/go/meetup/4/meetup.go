package meetup

import "time"

type WeekSchedule int

const (
    First WeekSchedule = iota
    Second
    Third
    Fourth
    Teenth
    Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
    var startDate time.Time
	switch wSched {
    case First:
        startDate = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC)
    case Second:
        startDate = time.Date(year, month, 8, 12, 0, 0, 0, time.UTC)
    case Third:
        startDate = time.Date(year, month, 15, 12, 0, 0, 0, time.UTC)
    case Fourth:
        startDate = time.Date(year, month, 22, 12, 0, 0, 0, time.UTC)
    case Teenth:
        startDate = time.Date(year, month, 13, 12, 0, 0, 0, time.UTC)
    case Last:
        startDate = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC).AddDate(0, 1, 0).AddDate(0, 0, -7)
    default:
        return 0
    }
    var t time.Time
    for i := 0; i < 7; i++ {
        if t = startDate.AddDate(0, 0, i); t.Weekday() == wDay {
            return t.Day()
        }
    }
    return 0
}