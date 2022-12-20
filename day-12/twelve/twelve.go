package twelve

import (
	"fmt"
	"math/rand"
)

type WalkerHist map[[2]int]bool

func (wh WalkerHist) Clone() WalkerHist {
	newHist := WalkerHist{}
	for k, v := range wh {
		newHist[k] = v
	}
	return newHist
}

type Walker struct {
	ID        int
	Ancestors []int
	Count     int
	History   WalkerHist
	Goal      Point
}

func (w *Walker) Walk(n *Node) []*Walker {
	walkers := []*Walker{}
	//
	if w.Visited(n) {
		fmt.Printf("visited node %v", n)
		return walkers
	}

	//fmt.Println(w.History)

	w.Visit(n)

	//fmt.Println("walker", w.ID, "walked to", n.Position)
	fmt.Println("walked to", n.Position)

	if n.Position.X == w.Goal.X && n.Position.Y == w.Goal.Y {
		fmt.Println("walker", w.ID, "reached goal in", w.Count)
		return walkers
	}

	// Try to walk to the right if a path is available
	if n.R != nil && !w.Visited(n.R) {
		walkers = append(walkers, w.Walk(n.R)...)
	}

	for _, i := range [3]*Node{n.D, n.L, n.U} {
		ww := w.Fork()
		if i == nil || ww.Visited(i) {
			continue
		}
		walkers = append(walkers, ww)
		walkers = append(walkers, ww.Walk(i)...)
	}

	return walkers
}

func (w *Walker) Fork() *Walker {
	ww := Walker{
		ID:      rand.Int(),
		Count:   w.Count,
		History: w.History.Clone(),
		Goal:    w.Goal,
	}
	ww.Ancestors = append(ww.Ancestors, w.ID)
	//fmt.Println(ww.History)
	return &ww
}

func (w *Walker) Visited(n *Node) bool {
	_, ok := w.History[n.Position.Coords()]
	return ok
}

func (w *Walker) Visit(n *Node) {
	w.History[n.Position.Coords()] = true
	w.Count++
}

type Point struct {
	X int
	Y int
}

func (p Point) Coords() [2]int {
	return [2]int{p.X, p.Y}
}

type Map struct {
	Grid  [][]int
	Start Point
	Goal  Point
}

type NodeMap struct {
	Grid  [][]*Node
	Start Point
	Goal  Point
}

type Node struct {
	Value    int
	Position Point
	U        *Node
	R        *Node
	D        *Node
	L        *Node
}

func (n *Node) Link(m NodeMap) {
	y := n.Position.Y
	x := n.Position.X
	traversible := func(o *Node) bool {
		return n.Value+1 >= o.Value
	}
	if y > 0 {
		up := m.Grid[y-1][x]
		if traversible(up) {
			n.U = up
		}
	}
	if y < len(m.Grid)-1 {
		down := m.Grid[y+1][x]
		if traversible(down) {
			n.D = down
		}
	}
	if x > 0 {
		left := m.Grid[y][x-1]
		if traversible(left) {
			n.L = left
		}
	}
	if x < len(m.Grid[y])-1 {
		right := m.Grid[y][x+1]
		if traversible(right) {
			n.R = right
		}
	}
}

type HeightMap map[rune]int

type NodeTree struct {
	Root *Node
}

func MakeHeightMap() HeightMap {
	m := HeightMap{}
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		m[r] = i
		i++
	}
	return m
}

func MakeMap(input []string, hm HeightMap) NodeMap {
	m := NodeMap{}

	for i := len(input) - 1; i >= 0; i-- {
		row := []*Node{}
		for j, k := range input[i] {
			var v int
			switch k {
			case 'S':
				v = hm['a']
				m.Start.Y = i
				m.Start.X = j
			case 'E':
				v = hm['z']
				m.Goal.Y = i
				m.Goal.X = j
			default:
				v = hm[k]
			}
			n := &Node{}
			n.Value = v
			n.Position.Y = i
			n.Position.X = j
			row = append(row, n)
		}
		m.Grid = append(m.Grid, row)
	}

	for _, i := range m.Grid {
		for _, n := range i {
			n.Link(m)
		}
	}
	return m
}
