package main

import (
	"fmt"
	"log"
	"lol/internal"
	"math"
	"sort"
	"strings"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()

	cardsums := make([]int, len(lines))
	for i, line := range lines {
		cards := strings.Split(line, ":")
		nums := strings.Split(cards[1], "|")
		left, right := make([]int, 10), make([]int, 25)
		fmt.Sscanf(nums[0], "%d %d %d %d %d %d %d %d %d %d", &left[0], &left[1], &left[2], &left[3], &left[4], &left[5], &left[6], &left[7], &left[8], &left[9])
		fmt.Sscanf(nums[1], "%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
			&right[0], &right[1], &right[2], &right[3], &right[4], &right[5], &right[6], &right[7], &right[8], &right[9],
			&right[10], &right[11], &right[12], &right[13], &right[14], &right[15], &right[16], &right[17], &right[18], &right[19],
			&right[20], &right[21], &right[22], &right[23], &right[24],
		)

		// sort all the numbers from the scratchcards for quicker processing
		sort.Ints(left)
		sort.Ints(right)
		matches := -1
		for _, e := range left {
			for _, s := range right {
				if e == s {
					matches += 1
					break
				} else if e < s {
					break
				}
			}
		}
		if matches > -1 {
			cardsums[i] = int(math.Pow(2, float64(matches)))
		}
	}

	// total all the cardsums
	var total int
	for _, a := range cardsums {
		total += a
	}
	log.Println(cardsums, total)
}
