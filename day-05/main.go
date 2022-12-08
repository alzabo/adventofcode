package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type collection struct {
	stacks map[int]*stack
}

func (c *collection) move(s, d int) error {
	for _, i := range []int{s, d} {
		_, ok := c.stacks[i]
		if !ok {
			return errors.New(fmt.Sprintf("no stack found with id %v", i))
		}
	}
	v, err := c.stacks[s].pop()
	if err != nil {
		return err
	}
	c.stacks[d].push(v)
	return nil
}

func (c *collection) inspect() {
	for k, v := range c.stacks {
		fmt.Println("id:", k, "items:", v)
	}
}

type stack struct {
	items []string
}

func (s *stack) push(i string) {
	s.items = append(s.items, i)
}

func (s *stack) pop() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("can't pop from empty stack")
	}
	i := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return i, nil
}

func moveItems(c collection, b []byte) error {
	moveExpr := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	moveMatch := moveExpr.FindSubmatch(b)
	if moveMatch == nil || len(moveMatch) != 4 {
		return errors.New("failed to parse move command")
	}
	// fmt.Printf("match: %q\n", moveMatch)
	amount, _ := strconv.Atoi(string(moveMatch[1]))
	src, _ := strconv.Atoi(string(moveMatch[2]))
	dst, _ := strconv.Atoi(string(moveMatch[3]))

	for i := 1; i <= amount; i++ {
		err := c.move(src, dst)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to execute move %s with message %v", b, err))
		}
	}
	return nil
}

func initialState(b [][]byte) collection {
	state := collection{}
	state.stacks = map[int]*stack{}

	for i := len(b) - 1; i >= 0; i-- {
		s := 1
		for j := 0; j < len(b[i]); s++ {
			bytes := b[i]

			// if no stack with this index exists in
			// the collection, create it
			_, ok := state.stacks[s]
			if !ok {
				state.stacks[s] = &stack{}
			} else {
				// each item value will be 3 characters
				// separated by a space. take 4 to consume
				// the space. the last item can't have 4
				// bytes sliced so just slice from the index
				// to the end
				var b []byte
				if (j + 4) > len(bytes) {
					b = bytes[j:]
				} else {
					b = bytes[j : j+4]
				}

				it := strings.TrimSpace(string(b))
				if it != "" {
					state.stacks[s].push(strings.TrimSpace(it))
				}
			}
			j += 4
		}
	}
	return state
}

func main() {
	i, _ := os.ReadFile("input")

	input := bytes.Split(i, []byte("\n"))
	var delim int

	// find the index of the empty line that separates
	// initial state from moves
	for i, l := range input {
		if len(l) == 0 {
			delim = i
		}
	}

	init := input[:delim]
	state := initialState(init)

	//state.inspect()

	for _, m := range input[delim:] {
		err := moveItems(state, m)
		if err != nil {
			fmt.Println("an error occurred when moving items...", err)
		}
		state.inspect()
	}

	topItems := []string{}
	for i := 1; i <= len(state.stacks); i++ {
		items := state.stacks[i].items
		topItems = append(topItems, strings.Trim(items[len(items)-1], "[]"))
	}
	fmt.Printf("top items: %s\n", strings.Join(topItems, ""))
}
