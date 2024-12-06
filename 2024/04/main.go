package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

type Offset struct {
	X int
	Y int
}

var part1Offsets = [][]Offset{
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},       // down
	{{0, 0}, {0, -1}, {0, -2}, {0, -3}},    // up
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},       // right
	{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},    // left
	{{0, 0}, {1, 1}, {2, 2}, {3, 3}},       // down, right
	{{0, 0}, {1, -1}, {2, -2}, {3, -3}},    // down, left
	{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}}, // up, left
	{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}},    // up, right
}

func part1() {
	count := 0
	b, _ := os.ReadFile("input.txt")
	rows := bytes.Split(b, []byte("\n"))
	for x, cols := range rows {
		for y, b := range cols {
			if b == 'X' {
				fmt.Printf("%d %d %s\n", x, y, string(b))
				for _, offsets := range part1Offsets {
					maybeXmas := make([]byte, 4)
					last := offsets[len(offsets)-1]
					xx := x + last.X
					yy := y + last.Y
					if xx < 0 || xx > len(cols)-1 || yy < 0 || yy > len(rows)-1 {
						continue
					}
					for i, o := range offsets {
						maybeXmas[i] = rows[x+o.X][y+o.Y]
					}
					if bytes.Equal(maybeXmas, []byte("XMAS")) {
						count += 1
					}
					//fmt.Println(last)
				}
			}
		}
	}
	fmt.Println("part1", count)
}
func part2() {
	count := 0
	xmasOffsets := [][]Offset{{{1, 1}, {-1, -1}}, {{-1, 1}, {1, -1}}}
	b, _ := os.ReadFile("input.txt")
	rows := bytes.Split(b, []byte("\n"))
	for x, cols := range rows {
		for y, b := range cols {
			if b != 'A' {
				continue
			}
			masCount := 0
			for _, o := range xmasOffsets {
				maybeMas := make([]byte, 3) // A match is AMS
				maybeMas[0] = 'A'
				x0 := x + o[0].X
				x1 := x + o[1].X
				y0 := y + o[0].Y
				y1 := y + o[1].Y

				// bounds check
				maxX := len(cols) - 1
				maxY := len(rows) - 1
				if x0 < 0 || x1 < 0 || x0 > maxX || x1 > maxX || y0 < 0 || y1 < 0 || y0 > maxY || y1 > maxY {
					continue
				}
				maybeMas[1] = rows[x0][y0]
				maybeMas[2] = rows[x1][y1]
				slices.Sort(maybeMas)
				fmt.Println(string(maybeMas))
				if slices.Equal(maybeMas, []byte("AMS")) {
					masCount += 1
				}
			}

			if masCount == 2 {
				count += 1
			}
		}
	}
	fmt.Println("part2", count)
}

func main() {
	part1()
	part2()
}
