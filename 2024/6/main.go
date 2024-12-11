package main

import (
	"aoc2024/internal"
	"log"
	"slices"
)

func main() {
	lines := internal.OpenInputFile("./input.txt").ReadLines()
	board_sz_x, board_sz_y := len(lines[0]), len(lines)
	board, g_start := getBoard(lines)

	// run one simulation to get the visited array
	visited, loop := simulateBoard(g_start, board, board_sz_x, board_sz_y)
	if loop {
		log.Fatal("Loop detected in first simulation")
	}

	// simulate guards movement by adding artificial obstacles
	total_possible_obstacles := 0
	for _, v := range visited {
		obs_a := location{v.x, v.y, OBSTACLE}
		new_board := append(slices.Clone(board), obs_a)

		_, loop := simulateBoard(g_start, new_board, board_sz_x, board_sz_y)
		if loop {
			total_possible_obstacles++
		}
	}

	// log.Println("Visited count:", len(visited), board_sz_x, board_sz_y)
	log.Println("Total possible obstacles:", total_possible_obstacles)
}

type object rune

const (
	OBSTACLE object = '#'
	VISITED  object = 'X'
)

type location struct {
	x, y int
	obj  object
}

type direction rune

const (
	UP    direction = '^'
	DOWN  direction = 'v'
	LEFT  direction = '<'
	RIGHT direction = '>'
)

func (d direction) String() string {
	return string(d)
}

// Turn returns the direction that the guard should turn to
func (d direction) Turn() direction {
	var res direction
	switch d {
	case UP:
		res = RIGHT
	case RIGHT:
		res = DOWN
	case DOWN:
		res = LEFT
	case LEFT:
		res = UP
	default:
		return d
	}
	return res
}

type guard struct {
	x, y int
	dir  direction
}

func (g guard) NextPossibleMove(board []location) guard {
	switch g.dir {
	case UP:
		if slices.Contains(board, location{g.x, g.y - 1, OBSTACLE}) {
			g.dir = g.dir.Turn()
		} else {
			g.y--
		}
	case RIGHT:
		if slices.Contains(board, location{g.x + 1, g.y, OBSTACLE}) {
			g.dir = g.dir.Turn()
		} else {
			g.x++
		}
	case DOWN:
		if slices.Contains(board, location{g.x, g.y + 1, OBSTACLE}) {
			g.dir = g.dir.Turn()
		} else {
			g.y++
		}
	case LEFT:
		if slices.Contains(board, location{g.x - 1, g.y, OBSTACLE}) {
			g.dir = g.dir.Turn()
		} else {
			g.x--
		}
	}
	return g
}

func getBoard(lines []string) ([]location, guard) {
	var board []location // board is a sparse matrix of locations (x, y, obj)
	var g guard
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '#':
				board = append(board, location{x, y, OBSTACLE})
			case '^':
				g = guard{x, y, UP}
			case '>':
				g = guard{x, y, RIGHT}
			case '<':
				g = guard{x, y, LEFT}
			case 'v':
				g = guard{x, y, DOWN}
			}
		}
	}
	return board, g
}

func simulateBoard(g guard, board []location, board_sz_x, board_sz_y int) ([]location, bool) {
	visited, guard_prev_pos := []location{}, []guard{}
	for { // simulate guards movement within the board
		if !slices.Contains(visited, location{g.x, g.y, VISITED}) { // already visited this cell
			visited = append(visited, location{g.x, g.y, VISITED})
		}
		g = g.NextPossibleMove(board)
		if g.x < 0 || g.x >= board_sz_x || g.y < 0 || g.y >= board_sz_y { // out of bounds
			break
		}
		if slices.Contains(guard_prev_pos, g) { // loop detected, arrived at starting point
			return nil, true
		}
		guard_prev_pos = append(guard_prev_pos, g) // add to the list of guard position history
	}
	return visited, false
}
