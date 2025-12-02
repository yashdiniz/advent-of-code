package main

import (
	"aoc2025/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	line := input.ReadLines()[0] // data will always be in a single line

	sum := 0
	for _, r := range parse(line) {
		for i := r[0]; i <= r[1]; i++ {
			v := strconv.Itoa(i)
			vlen := len(v)
			if vlen >= 2 && vlen%2 == 0 {
				l, r := v[:vlen/2], v[vlen/2:]
				if l == r {
					log.Print("Invalid found:", i)
					sum += i
				}
			}
		}
	}
	log.Print("Answer:", sum)
}

func parse(line string) [][2]int {
	var out [][2]int
	// split list by comma
	for v := range strings.SplitSeq(line, ",") {
		// then split each range by hyphen
		r := strings.Split(v, "-")
		start, err := strconv.Atoi(r[0])
		if err != nil {
			log.Panic("failed at parse for start")
		}
		end, err := strconv.Atoi(r[1])
		if err != nil {
			log.Panic("failed at parse for end")
		}
		out = append(out, [2]int{start, end})
	}
	return out
}
