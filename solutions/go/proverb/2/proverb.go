package proverb

import "fmt"

var (
    proverbFormat = "For want of a %v the %v was lost."
    endingFormat = "And all for the want of a %v."
)

func Proverb(rhyme []string) []string {
    proverbs := make([]string, len(rhyme))
    for i := 0; i < len(rhyme) - 1; i++ {
    	proverbs[i] = fmt.Sprintf(proverbFormat, rhyme[i], rhyme[i + 1])   
    }
    if len(rhyme) > 0 {
        proverbs[len(rhyme) - 1] = fmt.Sprintf(endingFormat, rhyme[0])
    }
    return proverbs
}