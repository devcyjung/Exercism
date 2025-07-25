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
    case Teenth:
        startDate = time.Date(year, month, 13, 12, 0, 0, 0, time.UTC)
    case Last:
        startDate = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC).AddDate(0, 1, 0).AddDate(0, 0, -7)
    case First:
        startDate = time.Date(year, month, 1 + 7 * 0, 12, 0, 0, 0, time.UTC)
    case Second:
        startDate = time.Date(year, month, 1 + 7 * 1, 12, 0, 0, 0, time.UTC)
    case Third:
        startDate = time.Date(year, month, 1 + 7 * 2, 12, 0, 0, 0, time.UTC)
    case Fourth:
        startDate = time.Date(year, month, 1 + 7 * 3, 12, 0, 0, 0, time.UTC)
    default:
        return -1
    }
    var t time.Time
    for i := 0; i < 7; i++ {
        if t = startDate.AddDate(0, 0, i); t.Weekday() == wDay {
            return t.Day()
        }
    }
    return -1
}