package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    hour, minute int
}

func getTime(v int) (int, int) {
    for v < 0 {
        v += 24 * 60  
    }
    v %= 24 * 60
    return v/60, v%60
}

func New(h, m int) Clock {
    hour, minute := getTime(60 * h + m)
	return Clock{
        hour,
        minute,
    }
}

func (c Clock) Add(m int) Clock {
    c.hour, c.minute = getTime(m + c.minute + 60 * c.hour)
    return c
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
