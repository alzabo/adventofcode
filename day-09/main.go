package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alzabo/adventofcode-2022/nine"
)

func main() {
	fmt.Println("Day 9")
	rope := nine.NewRope()
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := bytes.Split(b, []byte("\n"))
	nine.ExecuteMoves(&rope, input)

	fmt.Println(len(rope.Head.Visited))

}
