package twelve

import (
    "fmt"
    "strings"
)

const (
    verseBeginFmt = "On the %v day of Christmas my true love gave to me: "
    itemFmt = "%v %v, "
    singleItemFmt = "%v %v."
    lastItemFmt = "and %v %v."
)

var (
    items = []string{
        "Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying",
        "Swans-a-Swimming", "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming",
    }
    cardinals = []string{
        "a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
    }
    ordinals = []string{
        "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
    }
)

func Verse(i int) string {
	switch {
    case i < 1:
        i = 0
    case i > 12:
        i = 11
    default:
        i--
    }
    var b strings.Builder
    b.WriteString(fmt.Sprintf(verseBeginFmt, ordinals[i]))
    if i == 0 {
        b.WriteString(fmt.Sprintf(singleItemFmt, cardinals[0], items[0]))
        return b.String()
    }
    for index := i; index > 0; index-- {
        b.WriteString(fmt.Sprintf(itemFmt, cardinals[index], items[index]))
    }
    b.WriteString(fmt.Sprintf(lastItemFmt, cardinals[0], items[0]))
    return b.String()
}

func Song() string {
    var b strings.Builder
    for i := 1; i <= 12; i++ {
        b.WriteString(Verse(i))
        if i < 12 {
        	b.WriteRune('\n')   
        }
    }
    return b.String()
}
