package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func abs(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return n * -1
	}
	return n
}

func part1() {
	b, _ := os.ReadFile("input.txt")
	lines := bytes.Split(b, []byte("\n"))
	count := 0
	for _, line := range lines {
		//fmt.Println(string(line))
		values := bytes.Split(line, []byte(" "))
		// delta absolute values must be between 1 and 3 inclusive
		//fmt.Println("delta of", last, level, ":", delta)
		// set first delta
		//fmt.Println(level)
		safe := check(values)
		if safe {
			count += 1
		}
	}
	fmt.Println("part1:", count)
}

func check(values [][]byte) bool {
	var last int
	var lastDelta int
	for i, val := range values {
		level, _ := strconv.Atoi(string(val))
		if i == 0 {
			last = level
			continue
		}
		delta := last - level

		if delta == 0 || abs(delta) > 3 {
			return false
		}
		last = level

		if i == 1 {
			lastDelta = delta
			continue
		}
		if !ok(delta, lastDelta) {
			return false
		}
		lastDelta = delta

		if i == len(values)-1 {
			return true
		}
	}
	return false
}

func part2() {
	b, _ := os.ReadFile("input.txt")
	lines := bytes.Split(b, []byte("\n"))
	count := 0
	for _, line := range lines {
		fmt.Println("full line")
		fmt.Println(string(line))
		values := bytes.Split(line, []byte(" "))
		safe := check(values)
		fmt.Println("safe:", safe)
		if safe {
			count += 1
			continue
		}
		for i := range len(values) {
			v := slices.Clone(values)
			v = slices.Delete(v, i, i+1)
			fmt.Println(string(bytes.Join(v, []byte(" "))))
			safe := check(v)
			fmt.Println("safe:", safe)
			if safe {
				count += 1
				break
			}
		}
	}
	fmt.Println("part2:", count)
}
func ok(a, b int) bool {
	if a > 0 && b > 0 {
		return true
	}
	if a < 0 && b < 0 {
		return true
	}
	return false
}

func main() {
	part1()
	part2()
}
