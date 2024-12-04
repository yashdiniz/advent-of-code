package main

import (
	"aoc2024/internal"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	line := strings.Join(internal.OpenInputFile("./input.txt").ReadLines(), "\n")

	// sum of products of all numbers, this time parse only returns enabled operations
	total := 0
	for _, token := range parse(line) {
		total += token.op1 * token.op2
	}
	log.Println("Final total is", total)
}

type token struct {
	op1 int
	op2 int
}

func parse(line string) []token {
	f_re := regexp.MustCompile(`(mul\(\d+,\d+\))|((do|don\'t)\(\))`)
	f_matches := f_re.FindAllString(line, -1)

	enabled := true // condition for enabling future multiplications
	result := []token{}
	for _, f_match := range f_matches {
		switch f_match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled { // skip disabled operations
				continue
			}

			re := regexp.MustCompile(`mul\((?P<op1>\d+),(?P<op2>\d+)\)`)
			match := re.FindStringSubmatch(f_match)
			op1_idx := re.SubexpIndex("op1")
			op2_idx := re.SubexpIndex("op2")

			op1, err := strconv.ParseInt(match[op1_idx], 10, 32)
			if err != nil {
				log.Fatal("failed in parse", err)
			}
			op2, err := strconv.ParseInt(match[op2_idx], 10, 32)
			if err != nil {
				log.Fatal("failed in parse", err)
			}

			// only append parsed results if followed by a do(), disable if followed by don't()
			// thus, it only returns enabled operations
			result = append(result, token{int(op1), int(op2)})
		}
	}

	return result
}
