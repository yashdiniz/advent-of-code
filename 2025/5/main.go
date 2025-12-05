package main

import (
	"aoc2025/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// parse part 1
	var ranges [][2]int
	var vals int
	for i, line := range lines {
		if line == "" {
			vals = i + 1
			break
		}
		d := strings.Split(line, "-")
		start, err := strconv.Atoi(d[0])
		if err != nil {
			log.Panic("failed in parse")
		}
		end, err := strconv.Atoi(d[1])
		if err != nil {
			log.Panic("failed in parse")
		}
		ranges = append(ranges, [2]int{start, end})
	}

	res := 0
	for _, line := range lines[vals:] {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Panic("failed in parse")
		}
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				res++
				break
			}
		}
	}
	log.Print("Answer:", res)
}
