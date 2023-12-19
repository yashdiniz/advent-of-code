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
			max_blue := 0
			for _, blue := range game.blue {
				if blue > max_blue {
					max_blue = blue
				}
			}
			max_red := 0
			for _, red := range game.red {
				if red > max_red {
					max_red = red
				}
			}
			max_green := 0
			for _, green := range game.green {
				if green > max_green {
					max_green = green
				}
			}
			result += max_blue * max_red * max_green
			log.Println(max_blue * max_red * max_green)
		}
		test_games()
	}
	log.Println("result", result)
}
