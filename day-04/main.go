package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func overlap(a, b string) bool {
	if strings.Contains(a, b) || strings.Contains(b, a) {
		return true
	}
	return false
}

func expands(s []byte) (string, error) {
	b, a, ok := bytes.Cut(s, []byte("-"))
	if !ok {
		return "", errors.New("could not parse section")
	}
	start, err := strconv.Atoi(string(b))
	if err != nil {
		return "", errors.New("failed to convert range start")
	}
	end, err := strconv.Atoi(string(a))
	if err != nil {
		return "", errors.New("failed to convert range end")
	}
	fmt.Println("start", start, "end", end)
	sections := []string{}
	for start <= end {
		// using strings.Contains may have been a mistake compared to
		// using sets. the numeric representation is padded to 2 characters
		// to prevent "1" from matching "11", etc.
		sections = append(sections, fmt.Sprintf("%02d", start))
		start++
	}

	sec := strings.Join(sections, ",")

	return sec, nil
}

func expandSections(s []byte) (string, string, error) {
	b, a, ok := bytes.Cut(s, []byte(","))
	if !ok {
		return "", "", errors.New("could not parse section descriptor")
	}
	secb, err := expands(b)
	if err != nil {
		return "", "", err
	}
	seca, err := expands(a)
	if err != nil {
		return "", "", err
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
