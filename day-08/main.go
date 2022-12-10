package main

import (
	"bytes"
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
	return t.visibleUp() && t.visibleRight() && t.visibleDown() && t.visibleLeft()
}

// TODO: metaprogramming for these??
func (t *tree) visibleUp() bool {
	n := t.up
	for {
		if n == nil {
			return true
		}
		if n.height > t.height {
			return false
		}
		n = n.up
	}
}

func (t *tree) visibleRight() bool {
	if t.right == nil {
		return true
	}

	if t.right.height < t.height {
		return true
	}

	return t.right.visibleRight()
}

func (t *tree) visibleDown() bool {
	if t.down == nil {
		return true
	}

	if t.down.height < t.height {
		return true
	}

	return t.down.visibleDown()
}

func (t *tree) visibleLeft() bool {
	if t.left == nil {
		return true
	}

	if t.left.height >= t.height {
		return false
	}

	return t.left.visibleLeft()
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
func (g *treeGrid) add(x, y int, t *tree) {
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
			fmt.Printf("tree (%d, %d): %v, %v\n", y, x, t, t.visible())
		}
	}
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
			g.add(x, y, &tree)
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

	//parseFS(lines, root)

	//fmt.Println("part 1:", solve(root))

}
