package expenses

import "fmt"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
    res := make([]Record, 0)
	for _, v := range in {
        if predicate(v) {
            res = append(res, v)
        }
    }
    return res
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
        return r.Day >= p.From && r.Day <= p.To
    }
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        return r.Category == c
    }
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    r := 0.0
	for _, v := range Filter(in, ByDaysPeriod(p)) {
        r += v.Amount
    }
    return r
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	carr := Filter(in, ByCategory(c))
    if len(carr) == 0 {
        return 0, fmt.Errorf("unknown category %v", c)
    }
    return TotalByPeriod(carr, p), nil
}
