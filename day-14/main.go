package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Air  = iota
	Sand = iota
	Rock = iota
)

type Space int

type Reservoir struct {
	Grid [175][600]Space
}

func (r *Reservoir) AddRocks(s [][]byte) {
	rocks, err := parseRocks(s)
	if err != nil {
		fmt.Println("failed to get rock positions from input with error", err)
	}
	for _, pR := range rocks {
		pA := pR[0]
		for _, pB := range pR[1:] {
			xA := pA[0]
			xB := pB[0]
			yA := pA[1]
			yB := pB[1]

			if xA > xB {
				for x := xA; x >= xB; x-- {
					r.Grid[yA][x] = Rock
				}
			}
			if xA < xB {
				for x := xA; x <= xB; x++ {
					r.Grid[yA][x] = Rock
				}
			}
			if yA > yB {
				for y := yA; y >= yB; y-- {
					r.Grid[y][xA] = Rock
				}
			}
			if yA < yB {
				for y := yA; y <= yB; y++ {
					r.Grid[y][xA] = Rock
				}
			}
			pA = pB
		}
	}
}

func (r Reservoir) String() string {
	rows := []string{"## Visualization starts at column 480"}
	for y, i := range r.Grid {
		r := []string{}
		for x, j := range i {
			if x < 480 {
				continue
			}
			if y == 0 && x == 500 {
				r = append(r, "+")
				continue
			}
			switch j {
			case Air:
				r = append(r, ".")
			case Rock:
				r = append(r, "#")
			case Sand:
				r = append(r, "o")
			}
		}
		rows = append(rows, strings.Join(r, ""))
	}
	return strings.Join(rows, "\n")
}

type Point [2]int

func dropSand1(r *Reservoir) int {
	source := Point{500, 0}
	sandCount := 0
	sX := source[0]
	sY := source[1]

	for {
		fmt.Println("sX", sX, "sY", sY, r.Grid[sY][sX] == Sand)
		fmt.Println("sX", sX, "sY+1", sY+1, "Air:", r.Grid[sY+1][sX] == Air)

		// Out of bounds?
		if sY >= len(r.Grid)-2 {
			break
		}

		if r.Grid[sY+1][sX] == Air {
			fmt.Println("Got here!")
			sY++ // "fall" without setting state in the grid
			continue
		}

		// can't fall straight down any further, test to see if we can roll
		if r.Grid[sY+1][sX-1] == Air {
			sY++
			sX--
			continue
		}

		if r.Grid[sY+1][sX+1] == Air {
			sY++
			sX++
			continue
		}

		// if we got to here, the sand has nowhere else to go
		r.Grid[sY][sX] = Sand

		sandCount++
		sX = source[0]
		sY = source[1]
	}
	return sandCount
}

func parseRocks(s [][]byte) ([][]Point, error) {
	ps := [][]Point{}
	for _, i := range s {
		p, err := parseRock(i)
		if err != nil {
			return ps, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// parseRock operates on a byte slice consisting of x,y coordinates separated
// by the sequence ` -> `. A slice of Point is returned
func parseRock(b []byte) ([]Point, error) {
	points := []Point{}
	for _, c := range strings.Split(string(b), " -> ") {
		x, y, ok := strings.Cut(c, ",")
		if !ok {
			return points, fmt.Errorf("failed to parse line %s; could not split %s into x, y coordinates\n", b, c)
		}

		pt := Point{}
		for i, v := range [2]string{x, y} {
			var err error
			pt[i], err = strconv.Atoi(v)
			if err != nil {
				return points, fmt.Errorf("failed to parse %s (line %s); %v could not be converted to int\n", c, b, v)
			}
		}
		points = append(points, pt)
	}
	return points, nil
}

func readInput(f string) [][]byte {
	b, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return bytes.Split(b, []byte("\n"))
}

func main() {
	input := readInput("input")

	res := Reservoir{}
	res.AddRocks(input)
	fmt.Println(res.Grid[0][0] == Air)
	grains := dropSand1(&res)
	fmt.Println(res)
	fmt.Println("part 1:", grains)
}
