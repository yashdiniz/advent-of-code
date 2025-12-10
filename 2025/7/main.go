package main

import (
	"aoc2025/internal"
	"log"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// convert to grid of runes
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, c := range line {
			grid[i][j] = c
		}
	}

	// find S
	s_j := 0
	for j, c := range grid[0] {
		if c == 'S' {
			s_j = j
		}
	}

	memo := make(map[[2]int]int)
	log.Print("answer: ", worlds(grid, 0, s_j, memo))
}

func worlds(grid [][]rune, i, j int, memo map[[2]int]int) int {
	if i == len(grid) {
		return 1
	}
	log.Println("Entered:", i, j, string(grid[i][j]), memo[[2]int{i, j}])
	if v, ok := memo[[2]int{i, j}]; ok {
		log.Println("memoized result:", i, j, v)
		return v
	}

	cell := grid[i][j]
	var tmp int
	switch cell {
	case 'S':
		grid[i+1][j] = '|'
		tmp = worlds(grid, i+1, j, memo)
	case '^':
		if grid[i-1][j] == '|' {
			grid[i][j-1] = '|'
			tmp = worlds(grid, i+1, j-1, memo)
			grid[i][j+1] = '|'
			tmp += worlds(grid, i+1, j+1, memo)
		}
	case '.':
		if i > 0 && grid[i-1][j] == '|' {
			grid[i][j] = '|'
			tmp = worlds(grid, i+1, j, memo)
		}
	case '|':
		tmp = worlds(grid, i+1, j, memo)
	}

	memo[[2]int{i, j}] = tmp
	return tmp
}
