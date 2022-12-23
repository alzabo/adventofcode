package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func load(p string) [][]byte {
	b, err := os.ReadFile(p)
	if err != nil {
		fmt.Println(err)
	}
	res := bytes.Split(b, []byte("\n\n"))
	return res

}

func pairInOrder(bt []byte) (bool, error) {
	l, r, ok := bytes.Cut(bt, []byte("\n"))
	if !ok {
		return false, errors.New("could not split pair descriptor")
	}
	pl := parse(l)
	pr := parse(r)
	//m := listExpr.FindAllSubmatch(a, -1)
	//fmt.Printf("%s\n", m)
	//fmt.Printf("%s\n", a)
	//fmt.Println(parse(a))
	//fmt.Printf("%v\n", pairs)
	cmp := compare(pl, pr)
	switch cmp {
	case -1:
		return true, nil
	case 1:
		return false, nil
	default:
		return false, errors.New(fmt.Sprintf("comparison inconclusive! left: %v; right: %v", pl, pr))
	}
}

func compare(l, r []any) int {
	for i, ll := range l {
		// right ran out of items, left is larger
		if len(r)-1 < i {
			return 1
		}
		rr := r[i]
		switch ll.(type) {
		case int:
			switch rr.(type) {
			case int:
				lli := ll.(int)
				rri := rr.(int)
				if lli == rri {
					continue
				}
				if lli < rri {
					return -1
				}
				if lli > rri {
					return 1
				}
			case []any:
				cmp := compare([]any{ll}, rr.([]any))
				if cmp == 0 {
					continue
				} else {
					return cmp
				}
			}
		case []any:
			var crr []any
			cll := ll.([]any)
			switch rr.(type) {
			case []any:
				crr = rr.([]any)
				crrlen := len(crr)
				clllen := len(cll)

				// if both are 0 len, continue to test the next item
				if clllen == 0 && crrlen == 0 {
					continue
				}
				if clllen == 0 && crrlen > 0 {
					return -1
				}
				if clllen > 0 && crrlen == 0 {
					return 1
				}
			case int:
				crr = []any{rr.(int)}
			}
			cmp := compare(cll, crr)
			if cmp == 0 {
				continue
			} else {
				return cmp
			}
		}
	}
	// lazily compare as string representations. some equal lists may occur
	// when recursively called on inner slices
	if fmt.Sprintf("%s", l) == fmt.Sprintf("%s", r) {
		return 0
	}
	// if we got here, the comparison logic may have missed something
	// or right bay be larger than left
	fmt.Printf("left: %s; right: %s\n", l, r)
	return -1
}

func parse(bt []byte) []any {
	data := []any{}
	stack := []*[]any{&data}
	values := []byte("")

	//fmt.Printf("unwrapped: %s\n", bt[1:len(bt)-1])

	// Trim the first `[` and last `]`. We already made an outer
	// slice
	for _, b := range bt[1 : len(bt)-1] {
		//fmt.Printf("%v\n", stack)
		cs := stack[len(stack)-1]
		switch b {
		case byte('['):
			stack = append(stack, &[]any{})
		case byte(']'):
			parseValues(&values, cs)
			// append the current slice to the previous slice
			ps := stack[len(stack)-2]
			*ps = append(*ps, *cs)

			// remove the current slice from the stack
			stack = stack[:len(stack)-1]
		case byte(','):
			parseValues(&values, cs)
		default:
			values = append(values, b)
		}
	}
	return data
}

func parseValues(b *[]byte, s *[]any) {
	bb := bytes.TrimLeft(*b, "")
	inner := string(bb)
	*b = []byte("") // reset the byte slice
	for _, i := range strings.Split(inner, ",") {
		v, err := strconv.Atoi(i)
		if err != nil {
			//fmt.Println(err)
			continue
		}
		*s = append(*s, v)
	}
}

func main() {
	input := load("input")
	//fmt.Printf("%v\n\n", input)

	sumInOrder := 0
	for i, j := range input {
		_ = i
		fmt.Println(i)
		ordered, err := pairInOrder(j)
		if err != nil {
			panic(fmt.Sprintf("found an error %v when comparing %s", err, j))
		}
		if ordered {
			sumInOrder += i + 1 // example index starts with 1
		}
	}
	fmt.Println("part 1:", sumInOrder)

	allPackets := []any{
		[]any{[]any{2}},
		[]any{[]any{6}},
	}
	for _, i := range input {
		b, a, _ := bytes.Cut(i, []byte("\n"))
		allPackets = append(allPackets, parse(a), parse(b))
	}
	sort.Slice(allPackets, func(i, j int) bool {
		return compare(allPackets[i].([]any), allPackets[j].([]any)) == -1
	})
	fmt.Println(allPackets[0])
	fmt.Println(allPackets[len(allPackets)-1])

	part2indices := []int{}
	for i, j := range allPackets {
		if compare(j.([]any), []any{[]any{2}}) == 0 {
			part2indices = append(part2indices, i+1)
		}
		if compare(j.([]any), []any{[]any{6}}) == 0 {
			part2indices = append(part2indices, i+1)
		}
	}

	if len(part2indices) != 2 {
		panic("should only have 2 indices for part 2")
	}

	fmt.Println("part 2:", part2indices[0]*part2indices[1])

}
