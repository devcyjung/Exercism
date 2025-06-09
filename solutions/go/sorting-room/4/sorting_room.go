package sorting

import (
    "fmt"
    "strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	if t, ok := fnb.(FancyNumber); ok {
        v, err := strconv.Atoi(t.Value())
        if err != nil {
            return 0
        }
        return v
    }
    return 0
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if t, ok := fnb.(FancyNumber); ok {
        v, err := strconv.ParseFloat(t.Value(), 64)
        if err != nil {
            return "This is a fancy box containing the number 0.0"
        }
        return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(v))
    }
    return "This is a fancy box containing the number 0.0"
}

func DescribeAnything(i any) string {
	switch t := i.(type) {
        case float64:
        	return DescribeNumber(t)
        case int:
        	return DescribeNumber(float64(t))
        case NumberBox:
        	return DescribeNumberBox(t)
        case FancyNumberBox:
        	return DescribeFancyNumberBox(t)
        default:
        	return "Return to sender"
    }
}
