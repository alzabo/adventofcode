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
	//for _, w := range walkers {
	//	if w.Count > 985 {

	//		fmt.Printf("%v\n", w)
	//	}
	//}
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
	input, err := readInput("sample")
	if err != nil {
		fmt.Println(err)
	}

	m := twelve.MakeMap(input, hm)
	//fmt.Println(m)
	//fmt.Printf("%v", m.Grid[m.Goal.Y][m.Goal.X])
	goal := m.Grid[m.Goal.Y][m.Goal.X]
	_ = goal
	start := m.Grid[m.Start.Y][m.Start.X]
	_ = start
	walk(start, goal)

	fmt.Printf("%v\n", m.Grid[0][2])
	fmt.Printf("%v\n", m.Grid[0][3])
	//walkSome()

}

func walkSome() {
	w := twelve.Walker{
		ID:      1,
		History: twelve.WalkerHist{},
	}

	nodes := []*twelve.Node{
		{Position: twelve.Point{X: 0, Y: 0}},
		{Position: twelve.Point{X: 1, Y: 0}},
		{Position: twelve.Point{X: 2, Y: 0}},
		{Position: twelve.Point{X: 3, Y: 0}},
		{Position: twelve.Point{X: 4, Y: 0}},
		{Position: twelve.Point{X: 5, Y: 0}},
		{Position: twelve.Point{X: 6, Y: 0}},
		{Position: twelve.Point{X: 7, Y: 0}},
		{Position: twelve.Point{X: 8, Y: 0}},
	}
	for i, n := range nodes {
		if i < len(nodes)-1 {
			n.R = nodes[i+1]
		}
	}

	//if !cmp.Equal(Point{8, 0})

	_ = nodes[0].R.R.R.R

	w.Walk(nodes[0])
}
