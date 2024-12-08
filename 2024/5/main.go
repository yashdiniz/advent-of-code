package main

import (
	"aoc2024/internal"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()
	var pages [][]int

	// parse the rules
	for _, line := range lines {
		if strings.Contains(line, "|") {
			parseRules(line)
		} else if line == "" {
			continue
		} else if strings.Contains(line, ",") {
			var nums []int
			res := strings.Split(line, ",")
			for _, s := range res {
				f, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal("In parsing pages", err)
				}
				nums = append(nums, f)
			}
			pages = append(pages, nums)
		}
	}

	var validPages [][]int
	for _, page := range pages {
		if isPageValid(page) {
			validPages = append(validPages, page)
		}
	}

	// sum of middle numbers in each valid page
	sum := 0
	for _, page := range validPages {
		sum += page[len(page)/2]
	}
	log.Println("Sum of middle page numbers from valid pages:", sum)
}

type pages []int

var invRules = map[int]pages{}

func parseRules(line string) {
	parts := strings.Split(line, "|")
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("In parsing rules", err)
	}
	b, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("In parsing rules", err)
	}
	// rules[a] = append(rules[a], b)
	invRules[b] = append(invRules[b], a)
}

func isPageValid(page []int) bool {
	valid := true
	for i, num := range page {
		for _, n := range invRules[num] {
			if valid {
				valid = !slices.Contains(page[i:], n)
			} else { // if the page is invalid, we can skip the rest of the checks
				return false
			}
		}
	}
	return valid
}
