package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Inv    []int
	Op     func(int) int
	Test   func(int) int
	Count  int
	Modulo int
}

type Monkeys []*Monkey

func (m Monkeys) Len() int {
	return len(m)
}

func (m Monkeys) Less(i, j int) bool {
	return m[i].Count < m[j].Count
}

func (m Monkeys) Swap(i, j int) {
	mm := m[i]
	m[i] = m[j]
	m[j] = mm
}

type MonkeyTest struct {
	mod int
	t   int
	f   int
}

func (mt *MonkeyTest) F(i int) int {
	if i%mt.mod == 0 {
		return mt.t
	}
	return mt.f
}

func monkeyMaker(b []byte) *Monkey {
	m := &Monkey{}
	lines := bytes.Split(b, []byte("\n"))
	for i, j := range lines {
		_, value, _ := bytes.Cut(j, []byte(": "))
		switch i {
		case 0: // Monkey number line
			continue
		case 1: // list of items
			m.Inv = []int{}
			for _, v := range strings.Split(string(value), ", ") {
				i, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println("failed to parse", v, err)
				}
				m.Inv = append(m.Inv, i)
			}
		case 2: // operation func
			m.Op = makeOp(value)
		case 3:
			l := [][]byte{
				value,      // value passed to mod operator
				lines[i+1], // monkey index if evenly divisible
				lines[i+2], // monkey index if not evenly divisible
			}
			v := [3]int{}
			for i, j := range l {
				var err error
				v[i], err = lastDigit(j)
				if err != nil {
					fmt.Println("failed to parse", j, "with error", err)
				}
			}
			//fmt.Println("test values:", v)
			m.Test = makeTest(v)
			m.Modulo = v[0]
			//m.Test = MonkeyTest{
			//	mod: v[0],
			//	t:   v[1],
			//	f:   v[2],
			//}
			break
		}
		//fmt.Printf("%s\n", j)
	}
	return m
}

func makeTest(v [3]int) func(int) int {
	f := func(i int) int {
		if i%v[0] == 0 {
			return v[1]
		}
		return v[2]
	}
	return f
}

func makeOp(b []byte) func(int) int {
	// line will be something like
	// value = old * old
	var opf func(int, int) int
	var valb *int

	s := strings.Split(string(b), " ")
	vb := s[4]
	if vb != "old" {
		i, _ := strconv.Atoi(vb)
		valb = &i
	}

	oper := s[3]
	switch oper {
	case "*":
		opf = func(i, j int) int {
			return i * j
		}
	case "+":
		opf = func(i, j int) int {
			return i + j
		}
	}

	return func(i int) int {
		if valb != nil {
			return opf(i, *valb)
		}
		// if the operand was the non-numeric string "old",
		// valb will be nil and i should be passed twice
		return opf(i, i)
	}
}

func lastDigit(b []byte) (int, error) {
	v := strings.Split(string(b), " ")
	i, err := strconv.Atoi(v[len(v)-1])
	return i, err
}

func main() {
	i, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := bytes.Split(i, []byte("\n\n")) // two newlines separate each object

	monkeys := Monkeys{}
	for _, i := range input {
		monkeys = append(monkeys, monkeyMaker(i))
	}

	fmt.Println(monkeys[0].Test(999))
	fmt.Println(monkeys[1].Test(999))
	//fmt.Println(f(99))

	fmt.Println(monkeys)

	denominator := 1
	for _, m := range monkeys {
		denominator *= m.Modulo
	}

	for i := 0; i < 10000; i++ {
		for i, m := range monkeys {
			_ = i
			for {
				if len(m.Inv) == 0 {
					break
				}
				//fmt.Println("monkey", i, "has", m.Inv)
				x := m.Inv[0]
				//fmt.Println("monkey", i, "examines", x)
				m.Inv = append(m.Inv[:0], m.Inv[1:]...)
				//fmt.Println("monkey", i, "has", m.Inv, "remaining...")
				x = m.Op(x)
				x %= denominator
				nm := monkeys[m.Test(x)]
				nm.Inv = append(nm.Inv, x)
				m.Count++
				//fmt.Println("item with value", x, "thrown by monkey", i, "to", nm)
			}
		}
	}

	sort.Sort(monkeys)
	fmt.Println("part 2", monkeys[monkeys.Len()-1].Count*monkeys[monkeys.Len()-2].Count)
}
