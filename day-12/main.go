package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alzabo/adventofcode-2022/twelve"
)

type Walker struct {
	Pos   [2]int
	Count int
}

func walkm(m twelve.Map, start [2]int) []Walker {
	seen := twelve.WalkerHist{}
	queue := []Walker{{m.Start, 0}}
	goals := []Walker{}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if _, ok := seen[i.Pos]; ok {
			continue
		}
		seen[i.Pos] = true

		if i.Pos == m.Goal {
			goals = append(goals, i)
			continue
		}

		x := i.Pos[0]
		y := i.Pos[1]
		v := m.Grid[y][x]

		traversible := func(o int) bool {
			//fmt.Println("other:", o, "this:", v, v <= 1)
			return o-v <= 1
		}

		if x > 0 {
			lx := x - 1
			lf := m.Grid[y][lx]
			if traversible(lf) {
				queue = append(queue, Walker{[2]int{lx, y}, i.Count + 1})
			}
		}
		if x < len(m.Grid[y])-1 {
			rx := x + 1
			rt := m.Grid[y][rx]
			if traversible(rt) {
				queue = append(queue, Walker{[2]int{rx, y}, i.Count + 1})
			}
		}
		if y > 0 {
			uy := y - 1
			up := m.Grid[uy][x]
			if traversible(up) {
				queue = append(queue, Walker{[2]int{x, uy}, i.Count + 1})
			}
		}
		if y < len(m.Grid)-1 {
			dy := y + 1
			dn := m.Grid[dy][x]
			if traversible(dn) {
				queue = append(queue, Walker{[2]int{x, dy}, i.Count + 1})
			}
		}
	}

	return goals
}

func readInput(f string) ([]string, error) {
	input := []string{}

	i, err := os.ReadFile(f)
	if err != nil {
		return input, fmt.Errorf("%w", err)
	}
	is := bytes.Split(i, []byte("\n"))

	for _, b := range is {
		input = append(input, string(b))
	}
	return input, nil
}

func main() {
	hm := twelve.MakeHeightMap()
	input, err := readInput("input")
	if err != nil {
		fmt.Println(err)
	}

	m := twelve.MakeMapGrid(input, hm)

	// part 1
	for _, i := range walkm(m, m.Start) {
		fmt.Printf("part 1: %v\n", i)
	}
}
