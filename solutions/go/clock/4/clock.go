package clock

import "fmt"

type Clock struct {
    hours, minutes int
}

func getTime(v int) (hours, minutes int) {
    v %= 24 * 60
    v += 24 * 60
    v %= 24 * 60
    hours, minutes = v/60, v%60
    return
}

func New(hours, minutes int) Clock {
    hours, minutes = getTime(60 * hours + minutes)
	return Clock{
        hours,
        minutes,
    }
}

func (c Clock) Add(minutes int) Clock {
    c.hours, c.minutes = getTime(minutes + c.minutes + 60 * c.hours)
    return c
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
