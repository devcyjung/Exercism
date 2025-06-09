package queenattack

import "errors"

var (
    ErrInvalidFormat = errors.New("Input format is invalid")
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
    if whitePosition == blackPosition {
        return false, ErrInvalidFormat
    }
	err := errors.Join(validatePosition(whitePosition), validatePosition(blackPosition))
    if err != nil {
        return false, err
    }
    rowDiff, colDiff := whitePosition[0] - blackPosition[0], whitePosition[1] - blackPosition[1]
    return rowDiff == 0 || colDiff == 0 || rowDiff + colDiff == 0 || rowDiff - colDiff == 0, nil
}

func validatePosition(position string) error {
    if len(position) != 2 || !('a' <= position[0] && position[0] <= 'h') || !('1' <= position[1] && position[1] <= '8') {
        return ErrInvalidFormat
    }
    return nil
}