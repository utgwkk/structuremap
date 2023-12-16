package structuremap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/utgwkk/structuremap/testdata"
)

func TestEncodeSuccess(t *testing.T) {
	testcases := []struct {
		name  string
		input any
		want  map[string]any
	}{
		{
			name: "nil",
			input: (*struct {
				A string
			})(nil),
			want: nil,
		},
		{
			name: "simple",
			input: struct {
				A string
				B int
				C []byte
			}{
				A: "aaa",
				B: 100,
				C: []byte("hello"),
			},
			want: map[string]any{
				"A": "aaa",
				"B": 100,
				"C": []byte("hello"),
			},
		},
		{
			name: "ignore unexported field",
			input: testdata.HasUnexportedField{
				A: "aaa",
			},
			want: map[string]any{
				"A": "aaa",
			},
		},
	}
	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.input)
			if err != nil {
				t.Fatalf("got an error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Value mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func ptrOf[T any](t T) *T { return &t }

func TestEncodeFailure(t *testing.T) {
	testcases := []struct {
		name  string
		input any
	}{
		{
			name:  "string",
			input: "aaa",
		},
		{
			name:  "int",
			input: 100,
		},
		{
			name:  "string pointer",
			input: ptrOf("aaa"),
		},
	}
	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := Encode(tt.input)
			if err == nil {
				t.Error("expected an error")
			}
		})
	}
}
