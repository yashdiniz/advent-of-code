package main

import (
	"aoc2024/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := internal.OpenInputFile("./test.txt").ReadLines()
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

	log.Println(pages)
	log.Println(rules)
	log.Println(invRules)
}

type pages []int

var rules = map[int]pages{}
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
	rules[a] = append(rules[a], b)
	invRules[b] = append(invRules[b], a)
}
