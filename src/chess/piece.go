package chess

type PieceMoves struct {
	ChessPiece string
	Moves      [][]int
	MaxMoves   int
}

func (p PieceMoves) Pawn() PieceMoves {
	p.ChessPiece = "pawn"
	p.Moves = [][]int{{1, 0}}
	p.MaxMoves = 1
	return p
}

func (p PieceMoves) King() PieceMoves {
	p.ChessPiece = "king"
	p.Moves = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	p.MaxMoves = 8
	return p
}

func (p PieceMoves) Queen() PieceMoves {
	p.ChessPiece = "queen"
	p.Moves = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	p.MaxMoves = 27
	return p
}
