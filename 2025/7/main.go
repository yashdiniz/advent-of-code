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

	count := 0
	for i, row := range grid {
		for j, cell := range row {
			switch cell {
			case 'S':
				grid[i+1][j] = '|'
			case '^':
				if grid[i-1][j] == '|' {
					grid[i][j-1], grid[i][j+1] = '|', '|'
					count++
				}
			case '.':
				if i > 0 && grid[i-1][j] == '|' {
					grid[i][j] = '|'
				}
			}
		}
	}

	// print grid
	for _, r := range grid {
		log.Print(string(r))
	}
	log.Print("answer: ", count)
}
