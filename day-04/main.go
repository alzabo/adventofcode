package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func overlap(a, b []string) bool {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				return true
			}
		}
	}
	return false
}

func expands(s []byte) ([]string, error) {
	b, a, ok := bytes.Cut(s, []byte("-"))
	sections := []string{}
	if !ok {
		return sections, errors.New("could not parse section")
	}
	start, err := strconv.Atoi(string(b))
	if err != nil {
		return sections, errors.New("failed to convert range start")
	}
	end, err := strconv.Atoi(string(a))
	if err != nil {
		return sections, errors.New("failed to convert range end")
	}
	fmt.Println("start", start, "end", end)
	for start <= end {
		sections = append(sections, fmt.Sprintf("%02d", start))
		start++
	}

	return sections, nil
}

func expandSections(s []byte) ([]string, []string, error) {
	b, a, ok := bytes.Cut(s, []byte(","))
	if !ok {
		return []string{}, []string{}, errors.New("could not parse section descriptor")
	}
	secb, err := expands(b)
	if err != nil {
		return []string{}, []string{}, err
	}
	seca, err := expands(a)
	if err != nil {
		return []string{}, []string{}, err
	}

	return secb, seca, nil
}

func main() {
	var score1 int
	//limit := 10
	//var score2 int

	input, _ := os.ReadFile("input")
	s := bytes.Split(input, []byte("\n"))
	for i, l := range s {
		_ = i
		//if i > limit {
		//	break
		//}

		s1, s2, err := expandSections(l)
		if err != nil {
			fmt.Println(err)
		}
		if overlap(s1, s2) {
			score1 += 1
		}
		fmt.Println("input:", string(l), "sections:", s1, s2, overlap(s1, s2))
		//score1 += priority[match]
		//fmt.Println(match)
	}

	fmt.Println("total score (part 1):", score1)
	//fmt.Println("total score (part 2)", score2)

}
