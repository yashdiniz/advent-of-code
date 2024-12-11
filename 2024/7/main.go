package main

import (
	"aoc2024/internal"
	"log"
	"math"
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
		pos := int(math.Pow(3, float64(len(v)-1)))
		for i := 0; i < pos; i++ {
			res := v[0]
			flags := numToBase3(i, len(v)-1)

			for i, opnd := range v[1:] {
				switch flags[i] {
				case 0: // add
					res += opnd
				case 1: // multiply
					res *= opnd
				case 2: // concatenate
					g, err := strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(opnd))
					if err != nil {
						panic(err)
					}
					res = g
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

func numToBase3(num int, width int) []int {
	flags := make([]int, width)
	for i := 0; i < width; i++ {
		flags[i] = num % 3
		num /= 3
	}
	return flags
}

func numToBase2(num int, width int) []bool {
	flags := make([]bool, width)
	for i := 0; i < width; i++ {
		if num&(1<<i) != 0 {
			flags[i] = true
		}
	}
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
