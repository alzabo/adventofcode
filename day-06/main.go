package main

import (
	"bytes"
	"fmt"
	"os"
)

type tracker struct {
	chars []rune
	max   int
}

// add a new byte to tracker.b, drop bytes over max from
// the beginning of b
func (t *tracker) add(r rune) {
	t.chars = append(t.chars, r)
	if len(t.chars) > t.max {
		// TODO: review https://github.com/golang/go/wiki/SliceTricks in detail
		t.chars = t.chars[1:]
	}
}

func (t *tracker) len() int {
	return len(t.chars)
}

func (t *tracker) full() bool {
	fmt.Println(len(t.chars), t.max)
	return len(t.chars) == t.max
}

func (t *tracker) uniq() bool {
	counter := map[rune]int{}
	for _, i := range t.chars {
		counter[i] += 1
		if counter[i] > 1 {
			return false
		}
	}
	fmt.Println(counter)
	return true
}

func newTracker(m int) tracker {
	t := tracker{}
	t.chars = []rune{}
	t.max = m
	return t
}

func scan(s string) int {
	counter := 0
	t := newTracker(14)
	for _, ss := range s {
		t.add(ss)
		fmt.Println(t)
		counter++
		if !t.full() {
			continue
		}
		if t.uniq() {
			break
		}
	}
	return counter
}

func main() {
	i, _ := os.ReadFile("input")
	i = bytes.TrimSpace(i)

	res := scan(string(i))

	fmt.Println("counted bytes:", res)
}
