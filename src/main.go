package main

import (
	"chess/src/chess"
	"fmt"
)

func main() {
	var piece, position string
	var pos chess.Position

	fmt.Println("Enter the chess piece (queen, pawn, king): ")
	fmt.Scanln(&piece)

	fmt.Println("Enter the position of this chess piece on the board (a1, b6): ")
	fmt.Scanln(&position)

	pos, err := chess.ParsePosition(position)
	if err != nil {
		panic(err)
	}

	moves, err := chess.GetValidMoves(piece, pos)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Possible moves for %s are: %v \n", piece, moves)
}
