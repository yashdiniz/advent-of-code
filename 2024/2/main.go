package main

import (
	"aoc2024/internal"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()

	// parse the numbers from the lines
	var safe_cnt int
	for _, line := range lines {
		nums := parse(line)

		// eliminate unsafe patterns
		var distance int
		for i := 1; i < len(nums); i++ {
			d := nums[i] - nums[i-1]
			// rule 1: levels are either all increasing or all decreasing
			if d > 0 && distance < 0 { // unsafe pattern: positive distance in decreasing pattern
				distance = 0 // reset distance
				break
			}
			if d < 0 && distance > 0 { // unsafe pattern: negative distance in increasing pattern
				distance = 0 // reset distance
				break
			}
			distance += d

			// rule 2: adjacent levels differ between 1 to 3.
			if f := math.Abs(float64(d)); f < 1 || f > 3 { // unsafe pattern: not in range
				distance = 0 // reset distance
				break
			}
		}

		if distance != 0 {
			safe_cnt++ // safe pattern found
		}
	}

	log.Println("Final safe count is", safe_cnt)
}

func parse(line string) []int {
	var out []int
	numbers := strings.Split(line, " ")
	for _, n := range numbers {
		if n == "" {
			continue
		}
		v, err := strconv.ParseInt(n, 10, 32)
		if err != nil {
			log.Fatal("failed in parse", err)
		}
		out = append(out, int(v))
	}
	return out
}