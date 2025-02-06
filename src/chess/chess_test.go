package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParsePosition(t *testing.T) {
	// arguments that go to the function that we are testing
	type args struct {
		pos string
	}
	// tests are the actual test cases
	tests := []struct {
		name    string // <- name of the test
		args    args   // <- arguments for that test
		want    Position
		wantErr bool
	}{
		{
			name: "piece at a8",
			args: args{
				pos: "a8",
			},
			want: Position{},
		},
		{
			name: "piece at h8",
			args: args{
				pos: "h8",
			},
			want: Position{Col: 7, Row: 0},
		},
		{
			name: "piece at g5",
			args: args{
				pos: "g5",
			},
			want: Position{Col: 6, Row: 3},
		},
		{
			name: "piece at i4",
			args: args{
				pos: "i4",
			},
			want:    Position{0, 0},
			wantErr: true,
		},
		{
			name: "piece at e",
			args: args{
				pos: "e",
			},
			want:    Position{0, 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePosition(tt.args.pos)
			if (err != nil) != tt.wantErr {
				t.Errorf("positionToIndices() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("positionToIndices() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParsePosition(b *testing.B) {
	var (
		position Position
		err      error
	)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		position, err = ParsePosition("d5")
		if err != nil {
			b.Fatal(err)
		}
	}
	_, _ = fmt.Fprintln(io.Discard, position)
}

func TestPosition_String(t *testing.T) {
	type fields struct {
		Col int
		Row int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "col: 0, row: 0 should return A8",
			fields: fields{
				Col: 0, Row: 0,
			},
			want: "A8",
		},
		{
			name: "col: 7, row: 0 should return H8",
			fields: fields{
				Col: 7, Row: 0,
			},
			want: "H8",
		},
		{
			name: "col: 6, row: 3 should return G5",
			fields: fields{
				Col: 6, Row: 3,
			},
			want: "G5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				Col: tt.fields.Col,
				Row: tt.fields.Row,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPosition_String(b *testing.B) {
	var posStr string
	position := Position{Col: 6, Row: 3}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		posStr = position.String()
	}
	_, _ = fmt.Fprintln(io.Discard, posStr)
}

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
