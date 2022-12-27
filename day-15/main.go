package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
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

func main() {
	v := parseValves(readInput("input"))
	//fmt.Printf("%v\n", v)
	//for _, vv := range v {
	//	fmt.Printf("%v\n", vv)
	//}
	calc(v)
}
