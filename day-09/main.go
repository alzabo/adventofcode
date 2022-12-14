package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alzabo/adventofcode-2022/nine"
)

func main() {
	fmt.Println("Day 9")
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := bytes.Split(b, []byte("\n"))

	rope := nine.NewLongRope(2)
	nine.ExecuteMoves(rope.Knots[0], input)

	fmt.Println("Head visited", len(rope.Knots[0].Visited))
	fmt.Println("Tail visited", len(rope.Knots[1].Visited))

	longrope := nine.NewLongRope(10)
	nine.ExecuteMoves(longrope.Knots[0], input)
	for i, k := range longrope.Knots {
		fmt.Println(i, k.Position)
	}
	fmt.Println("part 2 tail visited", len(longrope.Knots[9].Visited))

}
