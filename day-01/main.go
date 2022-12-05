package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input")
	s := bytes.Split(input, []byte("\n"))

	counts := []int{0}

	for _, i := range s {
		if len(i) == 0 {
			counts = append(counts, 0)
			continue
		}
		v, err := strconv.Atoi(string(i))
		if err != nil {
			fmt.Println("Failed to convert bytes to int")
		}

		idx := len(counts) - 1
		counts[idx] += v
		//fmt.Println(len(i))
	}

	most := 0

	for _, i := range counts {
		if i > most {
			most = i
		}
	}

	fmt.Println(most)

}
