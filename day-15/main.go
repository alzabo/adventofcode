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
	start := "AA"
	for _, i := range m[start].Connections {
		for _, j := range i.Connections {
			fmt.Println("2:", j)
			for _, k := range j.Connections {
				fmt.Println("3:", k)
			}
		}
	}
}

func main() {
	v := parseValves(readInput("input"))
	//fmt.Printf("%v\n", v)
	//for _, vv := range v {
	//	fmt.Printf("%v\n", vv)
	//}
	calc(v)
}
