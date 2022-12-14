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

func TestPointTouch(t *testing.T) {
	tests := []struct {
		p1   Point
		p2   Point
		want bool
	}{
		{Point{X: 0, Y: 0}, Point{X: 0, Y: 0}, true},
		{Point{X: 0, Y: 0}, Point{X: 1, Y: 1}, true},
		{Point{X: 1, Y: 1}, Point{X: 0, Y: 0}, true},
		{Point{X: -2, Y: 0}, Point{X: -1, Y: 0}, true},
		{Point{X: 0, Y: 0}, Point{X: 2, Y: 1}, false},
		{Point{X: -2, Y: 0}, Point{X: 0, Y: 0}, false},
		{Point{X: -2, Y: 0}, Point{X: -4, Y: 0}, false},
		{Point{X: 0, Y: 0}, Point{X: -176, Y: 279}, false},
	}

	for _, tt := range tests {
		got := tt.p1.Touch(tt.p2)
		diff := cmp.Diff(tt.want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestPointDelta(t *testing.T) {
	tests := []struct {
		p1   Point
		p2   Point
		want PointDelta
	}{
		{Point{0, 0}, Point{0, 0}, PointDelta{0, 0}},
		{Point{0, 0}, Point{1, 1}, PointDelta{-1, -1}},
		{Point{X: -2, Y: 0}, Point{X: -1, Y: 0}, PointDelta{X: -1, Y: 0}},
		{Point{X: 0, Y: 0}, Point{X: 2, Y: 1}, PointDelta{X: -2, Y: -1}},
		{Point{X: -2, Y: 0}, Point{X: 0, Y: 0}, PointDelta{X: -2, Y: 0}},
		{Point{X: -2, Y: 0}, Point{X: -4, Y: 0}, PointDelta{X: 2, Y: 0}},
		{Point{X: 0, Y: 0}, Point{X: -176, Y: 279}, PointDelta{X: 176, Y: -279}},
	}

	for _, tt := range tests {
		got := tt.p1.Delta(tt.p2)
		diff := cmp.Diff(tt.want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestLongRopeMove(t *testing.T) {
	tests := []struct {
		moves [][]byte
		want  []Point
	}{
		//{
		//	[][]byte{[]byte("R 4")},
		//	[]Point{
		//		{X: 4, Y: 0}, // H
		//		{X: 3, Y: 0}, // 1
		//		{X: 2, Y: 0}, // 2
		//		{X: 1, Y: 0}, // 3
		//		{},
		//		{},
		//		{},
		//		{},
		//		{},
		//		{},
		//	},
		//},
		//{
		//	[][]byte{[]byte("R 4"), []byte("U 4")},
		//	[]Point{
		//		{X: 4, Y: 4}, // H
		//		{X: 4, Y: 3}, // 1
		//		{X: 4, Y: 2}, // 2
		//		{X: 3, Y: 2}, // 3
		//		{X: 2, Y: 2}, // 4
		//		{X: 1, Y: 1}, // 5
		//		{X: 0, Y: 0}, // 6
		//		{X: 0, Y: 0}, // 7
		//		{X: 0, Y: 0}, // 8
		//		{X: 0, Y: 0}, // 9
		//	},
		//},
		{
			[][]byte{
				[]byte("R 4"),
				[]byte("U 4"),
				[]byte("L 3"),
				[]byte("D 1"),
				[]byte("R 4"),
				[]byte("D 1"),
				[]byte("L 5"),
				[]byte("R 2"),
			},
			[]Point{
				{X: 2, Y: 2}, // H
				{X: 1, Y: 2}, // 1
				{X: 2, Y: 2}, // 2
				{X: 3, Y: 2}, // 3
				{X: 2, Y: 2}, // 4
				{X: 1, Y: 1}, // 5
				{X: 0, Y: 0}, // 6
				{X: 0, Y: 0}, // 7
				{X: 0, Y: 0}, // 8
				{X: 0, Y: 0}, // 9

			},
		},
	}
	moves := [][]byte{
		[]byte("R 4"),
		[]byte("U 4"),
		[]byte("L 3"),
		[]byte("D 1"),
		[]byte("R 4"),
		[]byte("D 1"),
		[]byte("L 5"),
		[]byte("R 2"),
	}
	_ = moves

	for _, tt := range tests {
		r := NewLongRope(10)
		ExecuteMoves(r.Knots[0], tt.moves)
		got := []Point{}
		for i := 0; i < 10; i++ {
			got = append(got, r.Knots[i].Position)
		}

		diff := cmp.Diff(tt.want, got)
		if diff != "" {
			t.Errorf("TestLongRopeMove() mismatch (-want +got):\n%s", diff)
		}
	}
}
