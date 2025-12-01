package main

import (
	"aoc2025/internal"
	"log"
	"strconv"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// safe starts at 50, need to count number of times it hits 0
	cnt0, safeptr := 0, 50
	for _, l := range lines {
		dir := parse(l)
		for i := 0; i < dir[1]; i++ {
			if dir[0] < 0 {
				safeptr--
			} else {
				safeptr++
			}
			if safeptr < 0 {
				safeptr += 100
			} else if safeptr >= 100 {
				safeptr %= 100
			}
			if safeptr == 0 {
				cnt0++
			}
		}
	}
	log.Println("Answer:", cnt0, safeptr)
}

func parse(line string) [2]int {
	var out [2]int
	if line[0] == 'L' {
		out[0] = -1
	} else {
		out[0] = 1
	}
	v, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal("failed in parse", err)
	}
	out[1] = v
	return out
}
