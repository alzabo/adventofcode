package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alzabo/adventofcode-2022/trees"
)

func main() {
	i, _ := os.ReadFile("input")
	lines := bytes.Split(i, []byte("\n"))

	t := trees.NewTreeGridFromInput(lines)

	t.Inspect()

	visibleCount := 0
	for _, v := range t.VisibleTrees() {
		if v {
			visibleCount++
		}
	}

	fmt.Println("visible trees (part 1):", visibleCount)

	treeA, _ := t.Fetch(97, 96)
	fmt.Println("a (97, 96):", &treeA, treeA, "visible:", treeA.Visible())
	//fmt.Println("a.down:", &treeA.down, treeA.down, "visible:", treeA.down.Visible())
	//fmt.Println("a.down.down", treeA.down.down)

	treeB, _ := t.Fetch(98, 96)
	fmt.Println("b (98, 96):", &treeB, treeB, "visible:", treeB.Visible())

	//parseFS(lines, root)

	//fmt.Println("part 1:", solve(root))

}
