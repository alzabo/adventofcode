package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type move struct {
	val int
}

type outcome struct {
	val int
}

var (
	rock     = move{1}
	paper    = move{2}
	scissors = move{3}
	win      = outcome{6}
	lose     = outcome{0}
	draw     = outcome{3}
	unknown  = outcome{}
)

var lookup = map[string]move{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

func parse(g []byte) (move, move, error) {
	a, b, found := bytes.Cut(g, []byte(" "))
	if !found {
		return move{}, move{}, errors.New("could not parse game")
	}

	ma := lookup[string(a)]
	mb := lookup[string(b)]

	return ma, mb, nil
}

func cmp(a move, b move) (res outcome, err error) {
	res = unknown
	if a == b {
		res = draw
	}
	if a == rock && b == scissors {
		res = lose
	}
	if a == rock && b == paper {
		res = win
	}
	if a == paper && b == scissors {
		res = win
	}
	if a == paper && b == rock {
		res = lose
	}
	if a == scissors && b == paper {
		res = lose
	}
	if a == scissors && b == rock {
		res = win
	}
	if res == unknown {
		err = errors.New("Could not determine outcome")
	}
	return
}

func play1(g []byte) (int, error) {

	elf, me, err := parse(g)

	if err != nil {
		return 0, errors.New("failed to parse the game")
	}

	result, err := cmp(elf, me)
	if err != nil {
		err = errors.New("could not determine the outcome of the game")
	}

	return result.val + me.val, nil
}

func main() {
	score1 := 0
	input, _ := os.ReadFile("input")

	s := bytes.Split(input, []byte("\n"))
	for _, l := range s {
		s1, err := play1(l)
		if err != nil {
			fmt.Println("failed to get score for", l)
		}
		score1 += s1
	}

	fmt.Println("total score (part 1)", score1)

}
