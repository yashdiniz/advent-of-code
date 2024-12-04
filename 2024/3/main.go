package main

import (
	"aoc2024/internal"
	"log"
	"regexp"
	"strconv"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()

	// sum of products of all numbers
	total := 0
	for _, line := range lines {
		for _, num := range parse(line) {
			total += num[0] * num[1]
		}
	}
	log.Println("Final total is", total)
}

func parse(line string) [][]int {
	mul_re := regexp.MustCompile(`mul\((?P<op1>\d+),(?P<op2>\d+)\)`)
	matches := mul_re.FindAllStringSubmatch(line, -1)
	op1_idx := mul_re.SubexpIndex("op1")
	op2_idx := mul_re.SubexpIndex("op2")

	// parge the numbers and multiply them
	result := make([][]int, len(matches))
	for i, match := range matches {
		op1, err := strconv.ParseInt(match[op1_idx], 10, 32)
		if err != nil {
			log.Fatal("failed in parse", err)
		}
		op2, err := strconv.ParseInt(match[op2_idx], 10, 32)
		if err != nil {
			log.Fatal("failed in parse", err)
		}
		result[i] = []int{int(op1), int(op2)}
	}

	return result
}
