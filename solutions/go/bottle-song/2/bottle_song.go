package bottlesong

import "fmt"

const (
    verse1Fmt = "%s green bottles hanging on the wall,"
    verse1Single = "One green bottle hanging on the wall,"
    verse2 = "And if one green bottle should accidentally fall,"
    verse3Fmt = "There'll be %s green bottles hanging on the wall."
    verse3Single = "There'll be one green bottle hanging on the wall."
)
    
var (
    numbers = []string{
    	"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	}
    capitalNumbers = []string{
        "No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
    }
)

func Recite(startBottles, takeDown int) []string {
    startIdx := startBottles
    if startBottles > 10 {
        startBottles = 10
    }
    endIdx := startBottles - takeDown
    if endIdx < 0 {
        endIdx = 0
    }
    if endIdx > startIdx {
        endIdx = startIdx
    }
    result := make([]string, 0, 5 * (startIdx - endIdx))
	for i := startIdx; i != endIdx; i-- {
        result = append(result, verse1(i), verse1(i), verse2, verse3(i), "")
    }
    result = result[:len(result) - 1]
    return result
}

func verse1(n int) string {
    if n == 1 {
        return verse1Single
    }
    return fmt.Sprintf(verse1Fmt, capitalNumbers[n])
}

func verse3(n int) string {
    if n == 2 {
        return verse3Single
    }
    return fmt.Sprintf(verse3Fmt, numbers[n - 1])
}