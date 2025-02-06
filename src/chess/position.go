package chess

import (
	"errors"
	"unicode"
)

type Position struct {
	// Not to return more than 3 values from a method.
	Col int
	Row int
}

var ErrInvalidPosition = errors.New("invalid position")

func ParsePosition(position string) (Position, error) {
	if len(position) != 2 {
		return Position{}, ErrInvalidPosition
	}

	col := int(unicode.ToUpper(rune(position[0])) - 'A')

	//You will face the chess board - white. Hence, starting point will be from top.
	//Chess pieces are moving only for one player.
	row := int('8' - unicode.ToUpper(rune(position[1])))

	if (row < 0 || row >= 8) || (col < 0 || col >= 8) {
		return Position{}, ErrInvalidPosition
	}

	return Position{Col: col, Row: row}, nil
}

// This is implementation of interface that has the value of an instantiation of struct always.
func (p Position) String() string {
	//Sprintf consumes a lot of time, hence the most optimised alternative.
	buf := []byte{
		byte('A' + p.Col),
		byte('8' - p.Row),
	}
	return string(buf)
}

func (p Position) isValid() bool {
	return p.Row >= 0 && p.Row < 8 && p.Col >= 0 && p.Col < 8
}
