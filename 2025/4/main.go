package main

import (
	"aoc2025/internal"
	"log"
)

func main() {
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	sum := 0
	for i, line := range lines {
		for j, roll := range line {
			total := 0
			if roll == '@' {
				for _, d := range dirs {
					x, y := i+d[0], j+d[1]
					if x < 0 || x >= len(lines) || y < 0 || y >= len(line) {
						continue
					}
					if lines[x][y] == '@' {
						total++
					}
				}
				if total < 4 {
					sum++
				}
			}
		}
	}
	log.Println("Answer:", sum)
}
