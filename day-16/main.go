package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	valveExpr = regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.+)$`)
)

type Valve struct {
	ID          string
	Value       int
	Connections []*Valve
}

//func (v Valve) String() string {
//	return fmt.Sprintf("ID: %v Value: %v Connections: [%v]", v.ID, v.Value, v.Connections)
//}

type ValveMap map[string]*Valve

func parseValves(b [][]byte) ValveMap {
	valves := ValveMap{}
	for _, i := range b {
		m := valveExpr.FindSubmatch(i)
		if len(m) != 4 {
			fmt.Printf("failed to parse %s", i)
		}
		id := string(m[1])

		v, ok := valves[id]
		if !ok {
			v = &Valve{ID: id}
			valves[id] = v
		}

		val, err := strconv.Atoi(string(m[2]))
		if err != nil {
			fmt.Println("failed to convert", m[2], "to int")
		}
		v.Value = val

		links := strings.Split(string(m[3]), ", ")
		for _, l := range links {
			lv, ok := valves[l]
			if !ok {
				lv = &Valve{ID: l}
				valves[l] = lv
			}
			v.Connections = append(v.Connections, lv)
		}
	}
	return valves
}

func readInput(n string) [][]byte {
	i, err := os.ReadFile(n)
	if err != nil {
		panic(fmt.Sprintf("error occurred when reading file %s\n", n))
	}
	return bytes.Split(i, []byte("\n"))
}

func calc(m ValveMap) {
	st := []struct {
		valve *Valve
		count int
		time  int
		op    map[string]struct{}
		hist  map[string]int
	}{
		{m["AA"], 0, 30, map[string]struct{}{}, map[string]int{"AA": 0}},
	}

	scores := map[int]struct{}{}
	high := 0

	for len(st) > 0 {
		i := st[len(st)-1]
		st = st[:len(st)-1]

		if i.time < 0 {
			continue
		}

		// prune paths that won't yield the highest scores
		if i.count < 350 && i.time < 24 {
			continue
		}

		if i.time == 0 && i.count > 0 {
			_, ok := scores[i.count]
			if !ok && i.count > high {
				high = i.count
				fmt.Println("high:", high, "timer:", i.time, "scored:", i.count, "opened:", i.op, "visited:", i.hist)
			}
			scores[i.count] = struct{}{}
			continue
		}

		c := i.hist[i.valve.ID]
		if c > 1 {
			continue
		}
		i.hist[i.valve.ID] += 1

		_, ok := i.op[i.valve.ID]
		if i.valve.Value > 0 && !ok {
			i.op[i.valve.ID] = struct{}{}
			i.time -= 1
			i.count += i.valve.Value * i.time
		}

		for _, c := range i.valve.Connections {
			n := i
			n.count = i.count
			n.valve = c
			n.time = i.time - 1

			opCp := map[string]struct{}{}
			for k := range i.op {
				opCp[k] = struct{}{}
			}
			n.op = opCp

			histCp := map[string]int{}
			for k, v := range i.hist {
				histCp[k] = v
			}
			n.hist = histCp

			st = append(st, n)
		}
	}
	//fmt.Println(scores)
}

type ValveEdge struct {
	Start string
	End   string
	Value int
	Cost  int
}

type actorState struct {
	name string
	time int
	pos  string
}

// https://www.reddit.com/r/adventofcode/comments/zn6k1l/comment/j0pewzt/?utm_source=share&utm_medium=web2x&context=3
// https://www.reddit.com/r/adventofcode/comments/zn6k1l/comment/j0rrokm/?utm_source=share&utm_medium=web2x&context=3
// https://en.wikipedia.org/wiki/Best-first_search
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// https://en.wikipedia.org/wiki/A*_search_algorithm
func calc2(m ValveMap) {
	filtered := ValveMap{"AA": m["AA"]}
	for _, v := range m {
		if v.Value > 0 {
			filtered[v.ID] = v
			fmt.Println(v)
		}
	}

	fmt.Println("")

	total := 0
	actors := []*actorState{
		{"Ele", 26, "AA"},
		{"Hum", 26, "AA"},
	}

	act := func(a *actorState) int {
		ve := []ValveEdge{}
		for _, v := range filtered {
			e := ValveEdge{
				Start: a.pos,
				End:   v.ID,
				Value: v.Value,
				Cost:  Dijkstra(filtered[a.pos], v),
			}
			if e.Value == 0 {
				continue
			}
			ve = append(ve, e)
		}

		// sort only by the score
		sort.Slice(ve, func(i, j int) bool {
			s1, _ := score(ve[i], a.time)
			s2, _ := score(ve[j], a.time)
			return s1 < s2
		})

		//sort.Slice(ve, func(i, j int) bool {
		//	return ve[i].Value-ve[i].Cost < ve[j].Value-ve[j].Cost
		//})

		// sort by cost, value
		sort.Slice(ve, func(i, j int) bool {
			if ve[i].Cost == ve[j].Cost {
				return ve[i].Value < ve[j].Value
			}
			return ve[i].Cost > ve[j].Cost
		})

		fmt.Println("actor", a, "moving...")

		for _, i := range ve {
			fmt.Println(i)
		}

		var toOpen ValveEdge
		var maxCost int
		for i := len(ve) - 1; i >= 0; i-- {
			if ve[i].Cost < a.time {
				if maxCost == 0 {
					//toOpen = ve[i]
					maxCost = ve[i].Cost + 1
				}
				if ve[i].Cost <= maxCost {
					if ve[i].Value-toOpen.Value > 1 {
						toOpen = ve[i]
					}
				}

				//if ve[i].Cost > 6 {
				//	continue
				//}

			}
		}

		if toOpen.Value > 0 {
			s, r := score(toOpen, a.time)
			fmt.Printf("%v opens valve %v for score %v; %v remains\n", a, toOpen.End, s, r)
			a.pos = toOpen.End
			filtered[toOpen.End].Value = 0
			a.time = r
			return s
		}
		fmt.Printf("%v could not open any valve\n", a)
		a.time = 0
		return 0
	}

	for {
		if actors[0].time == 0 && actors[1].time == 0 {
			break
		}

		// The actor with the most time remaining will act first.
		// this was something someone mentioned in a comment on reddit
		sort.Slice(actors, func(i, j int) bool {
			return actors[i].time < actors[j].time
		})

		if actors[1].time > 0 {
			total += act(actors[1])
		}
		if actors[0].time > 0 {
			total += act(actors[0])
		}
	}

	fmt.Println("total", total)

}

func score(ve ValveEdge, t int) (score int, remain int) {
	remain = t - ve.Cost
	score = ve.Value * remain
	return
}

type DijkstraQItem struct {
	Valve *Valve
	Prev  *DijkstraQItem
	Cost  int
}

// Because the cost of each path is always 1, we don't have to worry about finding
// a lower cost route or tracking the previous vertex
func Dijkstra(v1, v2 *Valve) int {
	q := []DijkstraQItem{{v1, nil, 0}}
	seen := map[string]bool{}

	for len(q) > 0 {
		qi := q[0]
		q = q[1:]

		seen[qi.Valve.ID] = true

		for _, c := range qi.Valve.Connections {
			if c.ID == v2.ID {
				return qi.Cost + 2
			}

			if _, ok := seen[c.ID]; ok {
				continue
			}

			n := DijkstraQItem{
				Valve: c,
				Prev:  nil,
				Cost:  qi.Cost + 1,
			}

			if len(q) == 0 {
				q = append(q, n)
			} else {
				idx := 0
				for i, qq := range q {
					idx = i
					if qq.Cost > n.Cost {
						break
					}
				}
				q = append(q[:idx], append([]DijkstraQItem{n}, q[idx:]...)...)
			}
		}
	}
	return 0
}

func main() {
	v := parseValves(readInput("input"))
	//fmt.Printf("%v\n", v)
	//for _, vv := range v {
	//	fmt.Printf("%v\n", vv)
	//}
	calc2(v)
}
