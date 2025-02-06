package chess

func GetValidMoves(piece string, position Position) ([]Position, error) {

	p, err := ParsePiece(piece)
	if err != nil {
		return nil, err
	}

	var moves = make([]Position, 0, p.MaxMoves)
	var tempPositions Position

	if p.ChessPiece == "queen" {
		for _, dir := range p.Moves {
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
		for _, dir := range p.Moves {
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
