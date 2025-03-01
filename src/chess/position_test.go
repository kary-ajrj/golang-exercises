package chess

import (
	"fmt"
	"io"
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
