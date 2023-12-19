package main

import (
	"log"
	"lol/internal"
	"regexp"
	"strconv"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()
	// lines := []string{"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	"......755.",
	// 	"...$.*....",
	// 	".664.598..",
	// }

	type number struct {
		num   int
		start int
		end   int
	}
	// map[line] = number data
	numbers := make([][]number, 140)

	// scan all the numbers in lines
	number_finder := regexp.MustCompile("[0-9]+")
	for i, line := range lines {
		inds := number_finder.FindAllStringSubmatchIndex(line, -1)
		nums := number_finder.FindAllString(line, -1)

		for j, ind := range inds {
			n, err := strconv.ParseInt(nums[j], 10, 64)
			if err != nil {
				log.Fatal("failed to parse numbers in line ", i, err)
			}
			num := number{
				num:   int(n),
				start: ind[0],
				end:   ind[1],
			}
			numbers[i] = append(numbers[i], num)
		}
	}

	engine_nums_sum := 0
	find_symbol := regexp.MustCompile("[^.0-9]")
	for i, line := range lines {
		// now for the 8-neighbour check
		for _, symbol := range find_symbol.FindAllStringIndex(line, -1) {
			ind := symbol[0]
			for _, num := range numbers[i] {
				if num.end == ind {
					engine_nums_sum += num.num
				}
				if num.start == ind+1 {
					engine_nums_sum += num.num
				}
			}
			for _, num := range numbers[i-1] {
				if ind >= num.start-1 && ind <= num.end {
					engine_nums_sum += num.num
				}
			}
			for _, num := range numbers[i+1] {
				if ind >= num.start-1 && ind <= num.end {
					engine_nums_sum += num.num
				}
			}
		}
	}
	log.Println(engine_nums_sum)
}
