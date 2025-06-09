package proverb

import "fmt"

func Proverb(rhyme []string) (r []string) {
    size := len(rhyme)
    if size == 1 {
        r = append(r, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    }
	for i, v := range rhyme {
        if i == 0 {
            continue
        }
        r = append(r, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], v))
        if i == size-1 {
            r = append(r, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
        }
    }
    return
}
