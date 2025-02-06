package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestGetValidMoves(t *testing.T) {
	type args struct {
		piece    string
		position string
	}
	tests := []struct {
		name                 string
		args                 args
		want                 []string
		wantErr              bool
		wantParsePositionErr bool
	}{
		{
			name: "queen at e4",
			args: args{position: "e4", piece: "queen"},
			want: []string{"D5", "C6", "B7", "A8", "E5", "E6", "E7", "E8", "F5", "G6", "H7", "D4", "C4", "B4", "A4", "F4", "G4", "H4", "D3", "C2", "B1", "E3", "E2", "E1", "F3", "G2", "H1"},
		},
		{
			name: "pawn at a2",
			args: args{position: "a2", piece: "pawn"},
			want: []string{"A1"},
		},
		{
			name: "king at d5",
			args: args{position: "d5", piece: "king"},
			want: []string{"C6", "D6", "E6", "C5", "E5", "C4", "D4", "E4"},
		},
		{
			name: "king at a1",
			args: args{position: "a1", piece: "king"},
			want: []string{"A2", "B2", "B1"},
		},
		{
			name: "queen at a8",
			args: args{position: "a8", piece: "queen"},
			want: []string{"B8", "C8", "D8", "E8", "F8", "G8", "H8", "A7", "A6", "A5", "A4", "A3", "A2", "A1", "B7", "C6", "D5", "E4", "F3", "G2", "H1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position := MustParsePosition(ParsePosition(tt.args.position))
			piece := MustParsePiece(ParsePiece(tt.args.piece))

			got, err := GetValidMoves(piece, position)
			gotString := PosToString(got)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetValidMoves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotString, tt.want) {
				t.Errorf("GetValidMoves() got = %v, want %v", gotString, tt.want)
			}
		})
	}
}

func MustParsePosition(position Position, err error) Position {
	if err != nil {
		panic(err)
	}
	return position
}

func MustParsePiece(piece Pieces, err error) Pieces {
	if err != nil {
		panic(err)
	}
	return piece
}

func PosToString(p []Position) []string {
	var moves = make([]string, 0, len(p))
	for _, pos := range p {
		moves = append(moves, pos.String())
	}
	return moves
}

func BenchmarkGetValidMoves(b *testing.B) {
	piece := "queen"
	p, err := ParsePiece(piece)
	if err != nil {
		b.Fatal(err)
	}

	position := Position{Col: 4, Row: 4}
	var result []Position

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ = GetValidMoves(p, position)
	}
	_, _ = fmt.Fprintln(io.Discard, result)
}
