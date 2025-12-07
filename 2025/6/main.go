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

	// parse all the lines into respective arrays
	ops_at := len(lines) - 1 // operators at the last line
	var ops []string
	opnds := make([][]int, ops_at)
	for i, line := range lines {
		if i == ops_at {
			for o := range strings.SplitSeq(line, " ") {
				if o == "" {
					continue
				}
				ops = append(ops, o)
			}
			break
		}
		for n := range strings.SplitSeq(line, " ") {
			if n == "" {
				continue
			}
			opnd, err := strconv.Atoi(n)
			if err != nil {
				log.Panic("failed at parse in main", err)
			}
			opnds[i] = append(opnds[i], opnd)
		}
	}

	// sum the results of each problem
	sum := 0
	for i, op := range ops {
		switch op {
		case "+":
			acc := 0
			for _, opnd := range opnds {
				acc += opnd[i]
			}
			sum += acc
		case "*":
			acc := 1
			for _, opnd := range opnds {
				acc *= opnd[i]
			}
			sum += acc
		}
	}

	log.Print("Answer: ", sum)
}
