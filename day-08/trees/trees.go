package trees

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type Tree struct {
	height int
	up     *Tree
	right  *Tree
	down   *Tree
	left   *Tree
}

func (t *Tree) Visible() bool {
	if t.up == nil || t.right == nil || t.down == nil || t.left == nil {
		return true
	}
	return t.visibleUp() || t.visibleRight() || t.visibleDown() || t.visibleLeft()
}

// TODO: metaprogramming for these??
func (t *Tree) visibleUp() bool {
	n := t.up
	for {
		if n == nil {
			return true
		}
		if n.height >= t.height {
			return false
		}
		// Follow edge to next ree in up direction
		n = n.up
	}
}

func (t *Tree) visibleRight() bool {
	n := t.right
	for {
		if n == nil {
			return true
		}
		if n.height >= t.height {
			return false
		}
		n = n.right
	}
}

func (t *Tree) visibleDown() bool {
	n := t.down
	for {
		if n == nil {
			return true
		}
		if n.height >= t.height {
			return false
		}
		n = n.down
	}
}

func (t *Tree) visibleLeft() bool {
	n := t.left
	for {
		if n == nil {
			return true
		}
		if n.height >= t.height {
			return false
		}
		n = n.left
	}
}

func (t *Tree) Score() int {
	c := map[string]int{
		"up":    0,
		"right": 0,
		"down":  0,
		"left":  0,
	}
	tUp := t.up
	for {
		if tUp == nil {
			break
		}
		c["up"]++
		if tUp.height >= t.height {
			break
		}
		tUp = tUp.up
	}

	tRight := t.right
	for {
		if tRight == nil {
			break
		}
		c["right"]++
		if tRight.height >= t.height {
			break
		}
		tRight = tRight.right
	}

	tDown := t.down
	for {
		if tDown == nil {
			break
		}
		c["down"]++
		if tDown.height >= t.height {
			break
		}
		tDown = tDown.down
	}

	tLeft := t.left
	for {
		if tLeft == nil {
			break
		}
		c["left"]++
		if tLeft.height >= t.height {
			break
		}
		tLeft = tLeft.left
	}
	return c["up"] * c["right"] * c["down"] * c["left"]
}

type TreeGrid struct {
	trees [][]*Tree
}

// add is called in to build a matrix of trees. it is assumed that
// it is called in order of ascending x, y coordinates
func (g *TreeGrid) add(y, x int, t *Tree) {
	// Add another ree slice to the y axis, if needed
	if len(g.trees)-1 < y {
		g.trees = append(g.trees, []*Tree{})
	}
	g.trees[y] = append(g.trees[y], t)
}

func (g *TreeGrid) linkTrees() {
	maxY := len(g.trees) - 1

	// for our purposes, the length of
	// each slice in the x axis will be
	// identical
	maxX := len(g.trees[maxY]) - 1

	for y, tt := range g.trees {
		for x, t := range tt {
			if x > 0 {
				t.left = g.trees[y][x-1]
			}
			if x < maxX {
				t.right = g.trees[y][x+1]
			}
			if y > 0 {
				t.up = g.trees[y-1][x]
			}
			if y < maxY {
				t.down = g.trees[y+1][x]
			}
		}
	}
}

func (g *TreeGrid) VisibleTrees() map[*Tree]bool {
	v := map[*Tree]bool{}
	for _, y := range g.trees {
		for _, t := range y {
			v[t] = t.Visible()
		}
	}
	return v
}

func (g *TreeGrid) HighScore() int {
	s := []int{}
	for _, y := range g.trees {
		for _, t := range y {
			s = append(s, t.Score())
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	//fmt.Println(s)

	return s[len(s)-1]
}

func (g *TreeGrid) Inspect() {
	for y, tt := range g.trees {
		for x, t := range tt {
			fmt.Printf("Tree %v (%d, %d): %v, %v\n", &t, y, x, t, t.Visible())
		}
	}
}

func (g *TreeGrid) Fetch(y, x int) (*Tree, error) {
	if y > len(g.trees)-1 {
		return &Tree{}, errors.New(fmt.Sprint("y index", y, "out of bounds"))
	}
	col := g.trees[y]

	if x > len(col)-1 {
		return &Tree{}, errors.New(fmt.Sprint("x index", x, "out of bounds"))
	}

	return g.trees[y][x], nil
}

func newTreeGrid() TreeGrid {
	g := TreeGrid{}
	g.trees = [][]*Tree{}
	return g
}

func NewTreeGridFromInput(b [][]byte) TreeGrid {
	g := newTreeGrid()
	for y, line := range b {
		_ = y
		for x, b := range line {
			_ = x
			v, _ := strconv.Atoi(string(b))
			tree := Tree{
				height: v,
			}
			g.add(y, x, &tree)
		}
	}
	g.linkTrees()
	return g
}
