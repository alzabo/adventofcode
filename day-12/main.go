package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"

	"github.com/alzabo/adventofcode-2022/twelve"
)

//func (w *Walker) Done() {
//
//}

func walk(n *twelve.Node, e *twelve.Node) {
	first := twelve.Walker{
		ID:        rand.Int(),
		Ancestors: []int{},
		Goal:      e.Position,
		History:   twelve.WalkerHist{},
	}
	walkers := []*twelve.Walker{&first}
	walkers = append(walkers, first.Walk(n)...)
	for _, w := range walkers {
		if w.Count > 985 {

			fmt.Printf("%v\n", w)
		}
	}
}

func readInput() ([]string, error) {
	input := []string{}

	i, err := os.ReadFile("input")
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
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	m := twelve.MakeMap(input, hm)
	//fmt.Println(m)
	//fmt.Printf("%v", m.Grid[m.Goal.Y][m.Goal.X])
	goal := m.Grid[m.Goal.Y][m.Goal.X]
	_ = goal
	start := m.Grid[m.Start.Y][m.Start.X]
	walk(start, goal)

}
