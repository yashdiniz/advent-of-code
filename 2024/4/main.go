package main

import (
	"aoc2024/internal"
	"log"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()
	line_num := len(lines)
	line_len := len(lines[0])

	// transpose input to check vertical
	lines_h := make([]string, line_len)
	for i := 0; i < line_num; i++ {
		for j := 0; j < line_len; j++ {
			lines_h[i] += string(lines[j][i])
		}
	}

	total := 0
	for i := 0; i < line_num; i++ {
		for j := 0; j < line_len-3; j++ {
			// horizontal check
			word := lines[i][j : j+4]
			if word == "XMAS" || word == "SAMX" {
				total++
			}
			// vertical check
			word = lines_h[i][j : j+4]
			if word == "XMAS" || word == "SAMX" {
				total++
			}
		}
	}
	// generate possible search strings, convolution-style
	for i := 0; i < line_num-3; i++ {
		for j := 0; j < line_len-3; j++ {
			// diagonal checks
			matrix := []string{
				lines[i][j : j+4],
				lines[i+1][j : j+4],
				lines[i+2][j : j+4],
				lines[i+3][j : j+4],
			}
			total += findXMASDiagonally(matrix)
		}
	}

	log.Println("XMAS count", total)
}

// count instances of XMAS in the matrix input
func findXMASDiagonally(input []string) int {
	total := 0
	// diagonal check
	word := ""
	for i := 0; i < 4; i++ { // LTR
		word += string(input[i][i])
	}
	if word == "XMAS" || word == "SAMX" {
		total++
	}
	word = ""
	for i := 0; i < 4; i++ { // RTL
		word += string(input[i][3-i])
	}
	if word == "XMAS" || word == "SAMX" {
		total++
	}

	return total
}
