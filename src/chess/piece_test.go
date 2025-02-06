package chess

import (
	"reflect"
	"testing"
)

func TestParsePiece(t *testing.T) {
	type args struct {
		piece string
	}
	tests := []struct {
		name    string
		args    args
		want    Pieces
		wantErr bool
	}{
		{
			name:    "google at a1",
			args:    args{piece: "google"},
			wantErr: true,
		},
		{
			name: "pawn at a",
			args: args{piece: "pawn"},
			want: Pieces{ChessPiece: "pawn", Moves: [][]int{{1, 0}}, MaxMoves: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePiece(tt.args.piece)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePiece() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePiece() got = %v, want %v", got, tt.want)
			}
		})
	}
}
