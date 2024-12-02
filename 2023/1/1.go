package main

import (
	"fmt"
	"log"
	"lol/internal"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	digits := []string{
		"0", "zero",
		"1", "one",
		"2", "two",
		"3", "three",
		"4", "four",
		"5", "five",
		"6", "six",
		"7", "seven",
		"8", "eight",
		"9", "nine",
	}
	var sum int64
	for _, line := range lines {
		first_index, last_index := len(line)+1, -1
		var first_digit, last_digit string
		for i, digit := range digits {
			if fi := strings.Index(line, digit); fi < first_index && fi != -1 {
				first_index = fi // find the left-most digit
				if i%2 == 1 {
					first_digit = digits[i-1]
				} else {
					first_digit = digit
				}
			}
			if li := strings.LastIndex(line, digit); li > last_index && li != -1 {
				last_index = li // find the right-most digit
				if i%2 == 1 {
					last_digit = digits[i-1]
				} else {
					last_digit = digit
				}
			}
		}

		calib_val := fmt.Sprintf("%v%v", first_digit, last_digit)
		val, err := strconv.ParseInt(calib_val, 10, 64)
		if err != nil {
			log.Fatal("failed in main ", err)
		}
		log.Printf("%v %v %v", calib_val, line, val)
		sum += val
	}
	log.Print("Final sum ", sum)
}
