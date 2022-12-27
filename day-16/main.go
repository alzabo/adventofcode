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

type actor struct {
	pos  *Valve
	hist map[string]int
}

type team struct {
	hum   actor
	ele   actor
	score int
	time  int
	op    map[string]struct{}
}

func calc2(m ValveMap) {
	st := []team{{
		hum:  actor{m["AA"], map[string]int{"AA": 0}},
		ele:  actor{m["AA"], map[string]int{"AA": 0}},
		op:   map[string]struct{}{},
		time: 26,
	}}

	high := 0

	for len(st) > 0 {
		i := st[0]
		st = st[1:]

		//fmt.Println("ok...", high, "hum", i.hum, "e", i.ele)

		if i.score > high {
			high = i.score
			fmt.Println("high", high, "team", i)
		}

		// prune paths that won't yield the highest scores
		if i.score < 300 && i.time < 22 {
			continue
		}

		if i.time < 0 {
			continue
		}

		//

		if c := i.ele.hist[i.ele.pos.ID]; c > 1 {
			continue
		}
		i.ele.hist[i.ele.pos.ID] += 1

		_, eok := i.op[i.ele.pos.ID]
		if i.ele.pos.Value > 0 && !eok {
			i.op[i.ele.pos.ID] = struct{}{}
			i.time -= 1
			i.score += i.ele.pos.Value * i.time
		}

		if c := i.hum.hist[i.hum.pos.ID]; c > 1 {
			continue
		}
		i.hum.hist[i.hum.pos.ID] += 1

		_, hok := i.op[i.hum.pos.ID]
		if i.hum.pos.Value > 0 && !hok {
			i.op[i.hum.pos.ID] = struct{}{}
			i.time -= 1
			i.score += i.hum.pos.Value * i.time
		}

		for _, ec := range i.ele.pos.Connections {
			for _, hc := range i.hum.pos.Connections {
				n := i
				n.ele.pos = ec
				n.hum.pos = hc
				n.time = i.time - 1

				//n.ele.hist = map[string]int{}
				//for k, v := range i.ele.hist {
				//	n.ele.hist[k] = v
				//}
				//n.hum.hist = map[string]int{}
				//for k, v := range i.hum.hist {
				//	n.hum.hist[k] = v
				//}
				st = append(st, n)
			}
		}
		//

		//

		//as := []actor{
		//	i.ele,
		//	i.hum,
		//}

		//for len(as) > 0 {
		//	a := as[len(as)-1]
		//	as = as[:len(as)-1]

		//	c := a.hist[a.pos.ID]
		//	if c > 1 {
		//		continue
		//	}
		//	a.hist[a.pos.ID] += 1

		//	_, ok := i.op[a.pos.ID]
		//	if a.pos.Value > 0 && !ok {
		//		i.op[a.pos.ID] = struct{}{}
		//		a.time -= 1
		//		i.score += a.pos.Value * a.time
		//	}

		//	for _, c := range a.pos.Connections {
		//		n := a
		//		n.pos = c
		//		n.time = a.time - 1

		//		histCp := map[string]int{}
		//		for k, v := range a.hist {
		//			histCp[k] = v
		//		}
		//		n.hist = histCp

		//		as = append(as, n)
		//	}

		//}

		//for _, a := range []*actor{&i.hum, &i.ele} {
	}
	//fmt.Println(scores)
}

func main() {
	v := parseValves(readInput("input"))
	//fmt.Printf("%v\n", v)
	//for _, vv := range v {
	//	fmt.Printf("%v\n", vv)
	//}
	calc2(v)
}
