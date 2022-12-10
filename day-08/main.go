package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	treeEdges
	height int
}

func (t *tree) visible() bool {
	if t.up == nil || t.right == nil || t.down == nil || t.left == nil {
		return true
	}
	return t.visibleUp() || t.visibleRight() || t.visibleDown() || t.visibleLeft()
}

// TODO: metaprogramming for these??
func (t *tree) visibleUp() bool {
	n := t.up
	for {
		if n == nil {
			return true
		}
		if n.height >= t.height {
			return false
		}
		// Follow edge to next tree in up direction
		n = n.up
	}
}

func (t *tree) visibleRight() bool {
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

func (t *tree) visibleDown() bool {
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

func (t *tree) visibleLeft() bool {
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

type treeEdges struct {
	up    *tree
	right *tree
	down  *tree
	left  *tree
}

type treeGrid struct {
	trees [][]*tree
}

// add is called in to build a matrix of trees. it is assumed that
// it is called in order of ascending x, y coordinates
func (g *treeGrid) add(y, x int, t *tree) {
	// Add another tree slice to the y axis, if needed
	if len(g.trees)-1 < y {
		g.trees = append(g.trees, []*tree{})
	}
	g.trees[y] = append(g.trees[y], t)
}

func (g *treeGrid) linkTrees() {
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

func (g *treeGrid) visibleTrees() map[*tree]bool {
	v := map[*tree]bool{}
	for _, y := range g.trees {
		for _, t := range y {
			v[t] = t.visible()
		}
	}
	return v
}

func (g *treeGrid) inspect() {
	for y, tt := range g.trees {
		for x, t := range tt {
			fmt.Printf("tree %v (%d, %d): %v, %v\n", &t, y, x, t, t.visible())
		}
	}
}

func (g *treeGrid) tree(y, x int) (*tree, error) {
	if y > len(g.trees)-1 {
		return &tree{}, errors.New(fmt.Sprint("y index", y, "out of bounds"))
	}
	col := g.trees[y]

	if x > len(col)-1 {
		return &tree{}, errors.New(fmt.Sprint("x index", x, "out of bounds"))
	}

	return g.trees[y][x], nil
}

func newTreeGrid() treeGrid {
	g := treeGrid{}
	g.trees = [][]*tree{}
	return g
}

func newTreeGridFromInput(b [][]byte) treeGrid {
	g := newTreeGrid()
	for y, line := range b {
		_ = y
		for x, b := range line {
			_ = x
			v, _ := strconv.Atoi(string(b))
			tree := tree{
				height: v,
			}
			g.add(y, x, &tree)
			//trees = append(trees, &tree)
		}
	}
	g.linkTrees()
	return g
}

func main() {
	i, _ := os.ReadFile("input")
	lines := bytes.Split(i, []byte("\n"))

	trees := newTreeGridFromInput(lines)

	trees.inspect()

	visibleCount := 0
	for _, v := range trees.visibleTrees() {
		if v {
			visibleCount++
		}
	}

	fmt.Println("visible trees (part 1):", visibleCount)

	treeA, _ := trees.tree(97, 96)
	fmt.Println("a (97, 96):", &treeA, treeA, "visible:", treeA.visible())
	fmt.Println("a.down:", &treeA.down, treeA.down, "visible:", treeA.down.visible())
	fmt.Println("a.down.down", treeA.down.down)

	treeB, _ := trees.tree(98, 96)
	fmt.Println("b (98, 96):", &treeB, treeB, "visible:", treeB.visible())

	//parseFS(lines, root)

	//fmt.Println("part 1:", solve(root))

}
