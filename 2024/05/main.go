package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	b, _ := os.ReadFile("input.txt")

	rules := [][2]int{}
	score := 0

line:
	for _, bb := range bytes.Split(b, []byte("\n")) {
		if bytes.ContainsRune(bb, '|') {
			a, b, ok := bytes.Cut(bb, []byte("|"))
			if !ok {
				panic("Failed to parse line")
			}
			//fmt.Println(string(bb))
			ai, _ := strconv.Atoi(string(a))
			bi, _ := strconv.Atoi(string(b))
			rules = append(rules, [2]int{ai, bi})
		} else if bytes.ContainsRune(bb, ',') {
			seq := []int{}
			for _, i := range bytes.Split(bb, []byte(",")) {
				s, _ := strconv.Atoi(string(i))
				seq = append(seq, s)
			}
			for i, s := range seq {
				if i == 0 {
					continue
				}
				for _, r := range rules {
					if r[0] == s {
						if slices.Contains(seq[:i], r[1]) {
							fmt.Println(s)
							continue line
						}
					}
				}
			}
			// in order
			fmt.Println(seq)
			middle := seq[len(seq)/2]
			score += middle
			fmt.Println(middle)
		}
	}
	fmt.Println("part1:", score)

}

func part2() {}
