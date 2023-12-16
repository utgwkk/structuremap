package structuremap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEncodeSuccess(t *testing.T) {
	testcases := []struct {
		name  string
		input any
		want  map[string]any
	}{
		{
			name: "nil",
			input: (*struct{
				A string
			})(nil),
			want: nil,
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
