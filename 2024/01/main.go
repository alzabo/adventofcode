package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func part1() {
	b, _ := os.ReadFile("input.txt")
	lines := bytes.Split(b, []byte("\n"))
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, l := range lines {
		s := bytes.Split(l, []byte(" "))
		//fmt.Println(s)
		left[i], _ = strconv.Atoi(string(s[0]))
		right[i], _ = strconv.Atoi(string(s[len(s)-1]))
	}

	slices.Sort(left)
	slices.Sort(right)
	dists := 0
	for i, v := range left {
		diff := abs(v - right[i])
		fmt.Printf("%d, %d: %d\n", v, right[i], diff)
		dists += diff
	}
	fmt.Println(dists)
}

func part2() {
	b, _ := os.ReadFile("input.txt")
	lines := bytes.Split(b, []byte("\n"))
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, l := range lines {
		s := bytes.Split(l, []byte(" "))
		//fmt.Println(s)
		left[i], _ = strconv.Atoi(string(s[0]))
		right[i], _ = strconv.Atoi(string(s[len(s)-1]))
	}

	score := 0
	for _, l := range left {
		count := 0
		for _, r := range right {
			if l == r {
				count += 1
			}
		}
		similar := l * count
		fmt.Println(l, count, similar)
		score += similar
	}
	fmt.Println(score)
}

func abs(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	part2()
}
