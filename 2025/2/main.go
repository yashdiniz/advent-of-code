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
			for j := 1; j < vlen; j++ {
				if vlen%j != 0 {
					continue
				}
				// string repeat check instead of split check
				if strings.Repeat(v[:j], vlen/j) == v {
					log.Print("gen:", strings.Repeat(v[:j], vlen/j), " v:", v, " vlen:", vlen, " j:", j)
					sum += i
					break
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
