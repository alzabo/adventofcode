package ten

import "fmt"

type Ten bool

type CPU struct {
	Cycle  int
	X      int
	Record []int
	Values map[int]int
	Screen *Screen
}

func (c *CPU) cycle() {
	c.Cycle++
	Draw(c.X, c.Cycle)
	for _, i := range c.Record {
		if i == c.Cycle {
			c.Values[c.Cycle] = c.Cycle * c.X
		}
	}
}

func (c *CPU) Noop() {
	c.cycle()
}

func (c *CPU) Addx(v int) {
	c.cycle()
	c.cycle()
	c.X += v
}

type Screen struct {
	Picture [240]string
}

func Draw(x int, c int) {
	lines := [][2]int{
		{1, 40},
		{41, 80},
		{81, 120},
		{121, 160},
		{161, 200},
		{201, 240},
	}
	//fmt.Println(c)
	for _, r := range lines {
		min := r[0]
		max := r[1]
		if c >= min && c <= max {
			if x >= c-min-1 && x <= c-min+1 {
				//fmt.Print(":", x, c-min, ":")
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		if c == max {
			fmt.Print("--\n")
		}
	}
}
