package sorting

import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
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

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
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

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if t, ok := fnb.(FancyNumber); ok {
        v, err := strconv.Atoi(t.Value())
        if err != nil {
            return "This is a fancy box containing the number 0.0"
        }
        return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(v))
    }
    return "This is a fancy box containing the number 0.0"
}

// DescribeAnything should return a string describing whatever it contains.
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
