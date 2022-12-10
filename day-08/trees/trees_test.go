package trees

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTreeScore(t *testing.T) {
	tests := []struct {
		tree Tree
		want int
	}{
		{Tree{}, 0},
		{
			Tree{
				height: 10,
				up:     &Tree{height: 1},
				right:  &Tree{height: 1},
				down:   &Tree{height: 1},
				left:   &Tree{height: 1},
			},
			1,
		},
		{
			Tree{
				height: 2,
				up:     &Tree{height: 1, up: &Tree{height: 1, up: &Tree{height: 1}}},
				right:  &Tree{height: 1},
				down:   &Tree{height: 1},
				left:   &Tree{height: 1},
			},
			3,
		},
		{
			Tree{
				height: 4,
				up:     &Tree{height: 1, up: &Tree{height: 4, up: &Tree{height: 1}}},
				right:  &Tree{height: 1},
				down:   &Tree{height: 1},
				left:   &Tree{height: 1, left: &Tree{height: 1}},
			},
			4,
		},
	}

	for _, tt := range tests {
		got := tt.tree.Score()
		diff := cmp.Diff(got, tt.want)
		if diff != "" {
			t.Fatalf(diff)

		}
	}

}
