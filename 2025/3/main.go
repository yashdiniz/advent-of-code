package main

import (
	"aoc2025/internal"
	"log"
	"strconv"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	sum := 0
	for _, line := range lines {
		max := 0
		for _, val := range combinations(line, 12) {

			num, err := strconv.Atoi(val)
			if err != nil {
				log.Panic("failed at main, in parse")
			}
			if num > max {
				max = num
			}
		}
		sum += max
		log.Print("max joltage: ", max)
	}
	log.Print("Answer: ", sum)
}

func combinations(line string, r int) []string {
	var result []string
	data := make([]rune, r) // Temporary slice to store the current combination
	for i := range data {
		data[i] = '0'
	}
	var combinationUtil func(int, int, []rune, *[]string, string) // Helper function for recursive backtracking

	combinationUtil = func(ind, r_size int, data []rune, result *[]string, line string) {
		if r_size == 0 {
			// A combination is found, append a copy to the result
			combination := make([]rune, len(data))
			copy(combination, data)
			*result = append(*result, string(combination))
			return
		}

		max := 0
		for i := ind; i <= len(line)-r_size; i++ {
			// Include the current element
			data[len(data)-r_size] = rune(line[i])
			num, err := strconv.Atoi(string(data))
			if err != nil {
				log.Panic("failed at combinations ", err)
			}
			if num > max { // eliminate entire paths if MSB is less than max discovered so far
				max = num
				combinationUtil(i+1, r_size-1, data, result, line) // Recurse for the next elements
			}
		}
	}

	combinationUtil(0, r, data, &result, line)
	return result
}
