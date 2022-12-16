package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alzabo/adventofcode-2022/ten"
)

func main() {
	t := ten.Ten(true)
	_ = t
	i, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := bytes.Split(i, []byte("\n"))

	cpu := ten.CPU{
		X:      1,
		Values: map[int]int{},
		Record: []int{
			20, 60, 100, 140, 180, 220,
		},
	}

	for _, i := range input {
		ins, val, _ := strings.Cut(string(i), " ")
		if ins == "noop" {
			cpu.Noop()
		}
		if ins == "addx" {
			v, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
			}
			cpu.Addx(v)
		}

	}

	fmt.Println(cpu.Values)
	pt1 := 0
	for _, v := range cpu.Values {
		pt1 += v
	}
	fmt.Println("part 1:", pt1)
}
