package main

import (
	"aoc2025/internal"
	"log"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()
	banks := parse(lines)

	sum := 0
	for _, bank := range banks {
		// var inds [2]int
		max := 0
		for i := range bank {
			for j := i + 1; j < len(bank); j++ {
				num := bank[i]*10 + bank[j]
				if num > max {
					max = num
					// inds = [2]int{i, j}
				}
			}
		}
		sum += max
		// log.Print("Bank: ", bank, " max_joltage: ", max, " inds: ", inds)
	}
	log.Print("Answer: ", sum)
}

func parse(lines []string) [][]int {
	var out [][]int = make([][]int, len(lines))
	for i, line := range lines {
		for _, digit := range line {
			out[i] = append(out[i], int(digit-0x30))
		}
	}
	return out
}
