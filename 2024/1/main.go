package main

import (
	"aoc2024/internal"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// parse the numbers from the lines
	var list_a, list_b []int
	for _, line := range lines {
		nums := parse(line)
		list_a = append(list_a, nums[0])
		list_b = append(list_b, nums[1])
	}

	// sort	the numbers into ascending order list_a and list_b
	sort.Slice(list_a, func(i, j int) bool {
		return list_a[i] < list_a[j]
	})
	// log.Print("list_a ", list_a)
	sort.Slice(list_b, func(i, j int) bool {
		return list_b[i] < list_b[j]
	})
	// log.Print("list_b ", list_b)

	// find the similarity between the two lists
	similarity := 0
	for _, a := range list_a {
		count := 0
		for _, b := range list_b {
			if a == b {
				count++
			}
		}
		similarity += count * a
	}
	log.Println("Final similarity is ", similarity)
}

func parse(line string) []int {
	var out []int
	numbers := strings.Split(line, " ")
	for _, n := range numbers {
		if n == "" {
			continue
		}
		v, err := strconv.ParseInt(n, 10, 32)
		if err != nil {
			log.Fatal("failed in parse", err)
		}
		out = append(out, int(v))
	}
	return out
}
