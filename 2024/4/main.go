package main

import (
	"aoc2024/internal"
	"log"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()
	line_num := len(lines)
	line_len := len(lines[0])

	total := 0
	// generate possible search strings, convolution-style
	for i := 0; i < line_num-2; i++ {
		for j := 0; j < line_len-2; j++ {
			// diagonal checks
			matrix := []string{
				lines[i][j : j+3],
				lines[i+1][j : j+3],
				lines[i+2][j : j+3],
			}
			total += findXMAS(matrix)
		}
	}

	log.Println("X-MAS count", total)
}

// count instances of X shaped MAS in the matrix input
func findXMAS(input []string) int {
	total := 0
	// diagonal check
	wordl, wordr := "", ""
	for i := 0; i < 3; i++ {
		wordl += string(input[i][i])
		wordr += string(input[i][2-i])
	}
	log.Println(wordl, wordr)
	if wordl == "MAS" && wordr == "SAM" {
		total++
	}
	if wordl == "MAS" && wordr == "MAS" {
		total++
	}
	if wordl == "SAM" && wordr == "SAM" {
		total++
	}
	if wordl == "SAM" && wordr == "MAS" {
		total++
	}

	return total
}
