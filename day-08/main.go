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

	tg := trees.NewTreeGridFromInput(lines)

	tg.Inspect()

	visibleCount := 0
	for _, v := range tg.VisibleTrees() {
		if v {
			visibleCount++
		}
	}

	fmt.Println("visible trees (part 1):", visibleCount)

	//treeA, _ := tg.Fetch(97, 96)
	//fmt.Println("a (97, 96):", &treeA, treeA, "visible:", treeA.Visible())
	//fmt.Println("a.down:", &treeA.down, treeA.down, "visible:", treeA.down.Visible())
	//fmt.Println("a.down.down", treeA.down.down)

	//treeB, _ := tg.Fetch(98, 96)
	//fmt.Println("b (98, 96):", &treeB, treeB, "visible:", treeB.Visible())
	fmt.Println("highest score (part 2):", tg.HighScore())

}
