package main

import (
	"aoc2024/internal"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()

	// parse input
	inputs := parseInput(lines)

	// calibration result
	cal_res := []int{}

	// add or multiply opnds such that the result is the same as the rhs
	for k, v := range inputs {
		// enumerate all possible combinations of operations
		pos := 1 << (len(v) - 1)
		for i := 0; i < pos; i++ {
			res := v[0]
			flags := bitMaskToFlags(i, len(v)-1)

			// add or multiply opnds such that the result is the same as the rhs
			for i, opnd := range v[1:] {
				if !flags[i] {
					res += opnd
				} else {
					res *= opnd
				}
			}

			if res == k && !slices.Contains(cal_res, k) {
				cal_res = append(cal_res, k)
			}
		}
	}

	sum := 0
	for _, r := range cal_res {
		sum += r
	}
	log.Println("Sum of all calibration results:", sum)
}

func bitMaskToFlags(num int, width int) []bool {
	flags := make([]bool, width)
	for i := 0; i < width; i++ {
		if num&(1<<i) != 0 {
			flags[i] = true
		}
	}
	slices.Reverse(flags)
	return flags
}

func parseInput(lines []string) map[int]([]int) {
	var res = make(map[int]([]int))
	for _, line := range lines {
		terms := strings.Split(line, ": ")
		opnds := strings.Split(terms[1], " ")
		rhs, err := strconv.Atoi(terms[0])
		if err != nil {
			panic(err)
		}
		if _, ok := res[rhs]; ok {
			panic("duplicate")
		}
		for _, opnd := range opnds {
			opnd, err := strconv.Atoi(opnd)
			if err != nil {
				panic(err)
			}
			res[rhs] = append(res[rhs], opnd)
		}
	}
	return res
}
