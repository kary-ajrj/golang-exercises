package chess

import (
	"errors"
	"fmt"
	"unicode"
)

var directions = map[string][][]int{
	"pawn": {{1, 0}},
	"king": {
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	},
	"queen": {
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	},
}

type Position struct {
	// Not to return more than 3 values from a method.
	Col int
	Row int
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

// This is implementation of interface that has the value of an instantiation of struct always.
func (p Position) isValid() bool {
	return p.Row >= 0 && p.Row < 8 && p.Col >= 0 && p.Col < 8
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

func GetValidMoves(piece string, position Position) ([]string, error) {
	var moves []string
	var tempPositions Position

	//You are accessing a map using the key.
	dirs, exists := directions[piece]

	if !exists {
		return nil, fmt.Errorf("invalid piece - %s", piece)
	}

	if piece == "queen" {
		for _, dir := range dirs {
			for step := 1; step < 8; step++ {
				tempPositions.Row = position.Row + dir[0]*step
				tempPositions.Col = position.Col + dir[1]*step
				if !tempPositions.isValid() {
					continue
				}
				moves = append(moves, tempPositions.String())
			}
		}
	} else {
		for _, dir := range dirs {
			tempPositions.Row = position.Row + dir[0]
			tempPositions.Col = position.Col + dir[1]
			if !tempPositions.isValid() {
				continue
			}
			moves = append(moves, tempPositions.String())
		}
	}
	return moves, nil
}
