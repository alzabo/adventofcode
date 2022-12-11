package nine

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHead(t *testing.T) {
	tests := []struct {
		head Head
		f    func(r *Head)
		want Head
	}{
		{NewHead(), func(h *Head) { h.Move(X, -2); h.Move(Y, 1) }, Head{
			Position: Point{X: -2, Y: 1},
			Visited:  map[[2]int]bool{{-2, 0}: true, {-1, 0}: true, {0, 0}: true, {-2, 1}: true},
		}},
	}

	for _, tt := range tests {
		tt.f(&tt.head)
		diff := cmp.Diff(tt.want, tt.head)
		if diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestRopeEndsTouching(t *testing.T) {
	tests := []struct {
		tail Tail
		head Head
		val  bool
		err  error
	}{
		{Tail{Position: Point{X: 0, Y: 0}}, Head{Position: Point{X: 0, Y: 0}}, true, nil},
		{Tail{Position: Point{X: 0, Y: 0}}, Head{Position: Point{X: 1, Y: 1}}, true, nil},
		{Tail{Position: Point{X: -2, Y: 0}}, Head{Position: Point{X: -1, Y: 0}}, true, nil},
		{Tail{Position: Point{X: 0, Y: 0}}, Head{Position: Point{X: 2, Y: 1}}, false, nil},
		{Tail{Position: Point{X: -2, Y: 0}}, Head{Position: Point{X: 0, Y: 0}}, false, nil},
	}

	for _, tt := range tests {
		tt.tail.Head = &tt.head
		v, err := tt.tail.EndsTouch()
		cmp.Equal(tt.err, err)
		cmp.Equal(tt.val, v)
	}
}

func TestNine(t *testing.T) {
	tests := []struct {
		input Nine
		want  bool
	}{{
		Nine(true), true,
	}}

	for _, tt := range tests {
		cmp.Equal(Nine(tt.want), tt.input)
	}
}
