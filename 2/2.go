package main

import (
	"fmt"
	"log"
	"lol/internal"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	type game struct {
		blue  []int
		red   []int
		green []int
		id    int
	}
	games := make([]game, len(lines))
	for i, line := range lines {
		id_split := strings.Split(line, ":")
		fmt.Sscanf(id_split[0], "Game %d", &games[i].id)
		for _, set := range strings.Split(id_split[1], ";") {
			for _, pull := range strings.Split(set, ",") {
				var count int
				var color string
				fmt.Sscanf(pull, "%d %s", &count, &color)
				switch color {
				case "blue":
					games[i].blue = append(games[i].blue, count)
				case "red":
					games[i].red = append(games[i].red, count)
				case "green":
					games[i].green = append(games[i].green, count)
				}
			}
		}
	}

	result := 0
	for _, game := range games {
		test_games := func() {
			for _, blue := range game.blue {
				if blue > 14 {
					return
				}
			}
			for _, red := range game.red {
				if red > 12 {
					return
				}
			}
			for _, green := range game.green {
				if green > 13 {
					return
				}
			}
			result += game.id
		}
		test_games()
	}
	log.Println("result", result)
}
