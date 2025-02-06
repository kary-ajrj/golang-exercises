package chess

import "errors"

type Pieces struct {
	ChessPiece string
	Moves      [][]int
	MaxMoves   int
}

func (p Pieces) Pawn() Pieces {
	p.ChessPiece = "pawn"
	p.Moves = [][]int{{1, 0}}
	p.MaxMoves = 1
	return p
}

func (p Pieces) King() Pieces {
	p.ChessPiece = "king"
	p.Moves = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	p.MaxMoves = 8
	return p
}

func (p Pieces) Queen() Pieces {
	p.ChessPiece = "queen"
	p.Moves = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	p.MaxMoves = 27
	return p
}

func ParsePiece(piece string) (Pieces, error) {
	p := Pieces{}
	switch piece {
	case "pawn":
		return p.Pawn(), nil
	case "king":
		return p.King(), nil
	case "queen":
		return p.Queen(), nil
	default:
		return Pieces{}, errors.New("invalid piece")
	}
}
