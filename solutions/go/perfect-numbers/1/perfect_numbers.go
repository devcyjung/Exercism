package perfect

import "errors"

type Classification int

const (
    ClassificationAbundant Classification = iota
    ClassificationDeficient
    ClassificationPerfect
    Invalid
)

var ErrOnlyPositive = errors.New("Must provide positive number")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
        return Invalid, ErrOnlyPositive
    }
    var aliquotSum int64
    for i := int64(1); i < n; i++ {
        if n % i == 0 {
            aliquotSum += i
        }
    }
    switch {
    case aliquotSum > n:
        return ClassificationAbundant, nil
    case aliquotSum < n:
        return ClassificationDeficient, nil
    default:
        return ClassificationPerfect, nil
    }
}