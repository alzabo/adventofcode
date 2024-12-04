package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	sum := 0
	f, _ := os.ReadFile("input.txt")
	mulExp, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	res := mulExp.FindAllSubmatch(f, -1)
	for _, r := range res {
		sum += mul(r)
		//fmt.Println(string(r[1]), string(r[2]))
	}
	fmt.Println("part1", sum)
}

func mul(r [][]byte) int {
	if len(r) != 3 {
		panic(fmt.Sprintf("Expected 3 elements, got %v", r))
	}
	a, _ := strconv.Atoi(string(r[1]))
	b, _ := strconv.Atoi(string(r[2]))
	return a * b
}

func part2() {
	sum := 0
	f, _ := os.ReadFile("input.txt")
	mulExp, _ := regexp.Compile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	res := mulExp.FindAllSubmatch(f, -1)
	do := true
	for _, r := range res {
		for _, b := range r {
			fmt.Print(string(b), " ")
		}
		if bytes.Equal(r[0], []byte("do()")) {
			do = true
		} else if bytes.Equal(r[0], []byte("don't()")) {
			do = false
		} else if bytes.Equal(r[0][:4], []byte("mul(")) {
			if do {
				sum += mul(r)
			}
		}
		fmt.Print("\n")
		//fmt.Println(string(r[1]), string(r[2]))
	}
	fmt.Println("part2", sum)
}

func main() {
	part1()
	part2()
}
