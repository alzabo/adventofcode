package twelve

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWalkerHistory(t *testing.T) {
	w1 := Walker{
		ID:      1,
		History: WalkerHist{},
	}

	w1.Visit(&Node{Position: Point{9, 9}})
	if !cmp.Equal(true, w1.Visited(&Node{Position: Point{9, 9}})) {
		t.Error("expected node to have been visited")
	}

	if !cmp.Equal(true, w1.Fork().Visited(&Node{Position: Point{9, 9}})) {
		t.Error("expected forked Walker to have visited node")
	}
}

func TestWalker(t *testing.T) {
	w := Walker{
		ID:      1,
		History: WalkerHist{},
		Goal:    Point{X: 8, Y: 0},
	}

	nodes := []*Node{
		{Position: Point{X: 0, Y: 2}},
		{Position: Point{X: 1, Y: 2}},
		{Position: Point{X: 2, Y: 2}},
		{Position: Point{X: 3, Y: 2}},
		{Position: Point{X: 4, Y: 2}},
		{Position: Point{X: 5, Y: 2}},
		{Position: Point{X: 6, Y: 2}},
		{Position: Point{X: 7, Y: 2}},
		{Position: Point{X: 8, Y: 2}},
	}
	for i, n := range nodes {
		if i < len(nodes)-1 {
			n.R = nodes[i+1]
		}
	}

	startNode := &Node{
		Position: Point{X: 0, Y: 1},
	}

	startNode.D = nodes[0]

	//if !cmp.Equal(Point{8, 0})

	_ = nodes[0].R.R.R.R

	ww := w.Walk(startNode)

	if !cmp.Equal(1, len(w.History)) {
		t.Error("expected 1 items in history, got", len(w.History), "history:", w.History)
	}

	if !cmp.Equal(10, len(ww[0].History)) {
		t.Error("expected 10 items in history, got", len(ww[0].History), "history:", ww[0].History)
	}
}
