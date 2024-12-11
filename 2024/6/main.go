package main

import (
	"aoc2024/internal"
	"log"
	"slices"
)

func main() {
	lines := internal.OpenInputFile("./test.txt").ReadLines()
	board_sz_x, board_sz_y := len(lines[0]), len(lines)
	board, g := getBoard(lines)
	var visited []location

	for {
		if !slices.Contains(visited, location{g.x, g.y, VISITED}) { // already visited this cell
			visited = append(visited, location{g.x, g.y, VISITED})
		}
		g = g.NextPossibleMove(board)
		if g.x < 0 || g.x >= board_sz_x || g.y < 0 || g.y >= board_sz_y { // out of bounds
			break
		}
	}

	log.Println("Visited count:", len(visited), board_sz_x, board_sz_y)
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
