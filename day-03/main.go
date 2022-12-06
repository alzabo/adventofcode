package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func findMatch(c []byte) (string, error) {
	match := ""
	cA := c[len(c)/2:]
	cB := c[:len(c)/2]
	setA := make(map[string]bool)

	for _, b := range cA {
		setA[string(b)] = true
	}
	for _, b := range cB {
		if setA[string(b)] {
			match = string(b)
			break
		}
	}

	if match == "" {
		return "", errors.New("Failed to find common item in contents")
	}

	return match, nil
}

func complement(a, b map[string]bool) map[string]bool {
	c := map[string]bool{}
	for k := range b {
		if a[k] {
			c[k] = true
		}
	}
	return c
}

func findBadge(s [][]byte) (string, error) {
	var sets []map[string]bool
	//fmt.Println("number of slices passed", len(s))
	for i, b := range s {
		sets = append(sets, map[string]bool{})
		for _, c := range b {
			char := string(c)
			fmt.Println(char)
			sets[i][char] = true
		}
	}

	//fmt.Println(sets)

	cmp := sets[0]
	for _, s := range sets[1:] {
		cmp = complement(cmp, s)
	}

	if len(cmp) < 1 {
		return "", errors.New("Could not find only 1 badge")
	}

	for i := range cmp {
		return i, nil
	}

	return "", errors.New("Could not find character with frequency in data set")
}

func main() {
	var score1 int
	var score2 int

	priority := make(map[string]int)
	v := 1

	// runes may be iterated as ints and may be cast to string
	for i := 'a'; i <= 'z'; i++ {
		priority[string(i)] = v
		v++
	}
	for i := 'A'; i <= 'Z'; i++ {
		priority[string(i)] = v
		v++
	}
	fmt.Println(priority)

	input, _ := os.ReadFile("input")
	s := bytes.Split(input, []byte("\n"))
	for _, l := range s {
		match, err := findMatch(l)
		if err != nil {
			fmt.Println("Failed to find match for", l)
		}
		score1 += priority[match]
		//fmt.Println(match)
	}

	i := 0
	for i < len(s) {
		j := i + 3
		fmt.Printf("slicing with %v:%v\n", i, j)
		sl := s[i:j]
		badge, err := findBadge(sl)
		if err != nil {
			fmt.Println("Couldn't find badge!")
		}
		score2 += priority[badge]
		i = j

		//fmt.Println("badge was", badge)
	}

	fmt.Println("total score (part 1)", score1)
	fmt.Println("total score (part 2)", score2)

}
