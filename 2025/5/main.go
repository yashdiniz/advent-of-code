package main

import (
	"aoc2025/internal"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// parse part 1
	var ranges [][2]int
	for _, line := range lines {
		if line == "" {
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

	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	log.Print(len(ranges))
	for i := 0; i < len(ranges); i++ {
		if i+1 == len(ranges) {
			break
		}
		item, next := ranges[i], ranges[i+1]
		if item == next { // delete duplicates
			ranges[i] = [2]int{0, 0}
			continue
		}
		// if the next item in the range is within current
		if next[0] >= item[0] && next[0] <= item[1] {
			if next[1] > item[1] {
				ranges[i+1] = [2]int{item[0], next[1]}
			} else {
				ranges[i+1] = item
			}
			ranges[i] = [2]int{0, 0}
		}
	}
	log.Print(len(ranges))

	sum := 0
	for _, v := range ranges {
		if v == [2]int{0, 0} {
			continue
		}
		log.Print(v)
		sum += v[1] - v[0] + 1
	}

	log.Print("Answer: ", sum)
}
