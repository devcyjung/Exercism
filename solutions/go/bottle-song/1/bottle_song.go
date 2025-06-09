package bottlesong

import "fmt"
    
var numbers = []string{
    "no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
}

func Recite(startBottles, takeDown int) (r []string) {
	for i := startBottles; i > startBottles-takeDown; i-- {
        if i != startBottles {
            r = append(r, "")
        }
        for j := 0; j < 2; j++ {
            r = append(r, fmt.Sprintf("%s green bottle%s hanging on the wall,", capitalize(numbers[i]), plurality(i)))
        }
        r = append(r, "And if one green bottle should accidentally fall,")
        r = append(r, fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", numbers[i-1], plurality(i-1)))
    }
    return
}

func capitalize(s string) string {
    runes := []rune(s)
    runes[0] += 'A' - 'a'
    return string(runes)
}

func plurality(s int) string {
    if s == 1 {
        return ""
    }
    return "s"
}