package chess

func GetValidMoves(piece Pieces, position Position) ([]Position, error) {

	var moves = make([]Position, 0, piece.MaxMoves)
	var tempPositions Position

	if piece.ChessPiece == "queen" {
		for _, dir := range piece.Moves {
			for step := 1; step < 8; step++ {
				tempPositions.Row = position.Row + dir[0]*step
				tempPositions.Col = position.Col + dir[1]*step
				if !tempPositions.isValid() {
					continue
				}
				moves = append(moves, tempPositions)
			}
		}
	} else {
		for _, dir := range piece.Moves {
			tempPositions.Row = position.Row + dir[0]
			tempPositions.Col = position.Col + dir[1]
			if !tempPositions.isValid() {
				continue
			}
			moves = append(moves, tempPositions)
		}
	}
	return moves, nil
}
