package main

import (
	"aoc2024/internal"
	"log"
)

func main() {
	lines := internal.OpenInputFile("./test.txt").ReadLines()
	board_sz_x, board_sz_y := len(lines[0]), len(lines)
	var board []location // board is a sparse matrix of locations (x, y, obj)

	// parse the board
	var g guard
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '#':
				board = append(board, location{x, y, obstacle})
			case '^':
				g = guard{x, y, up}
			case '>':
				g = guard{x, y, right}
			case '<':
				g = guard{x, y, left}
			case 'v':
				g = guard{x, y, down}
			}
		}
	}
	log.Println(board_sz_x, board_sz_y, board)
	log.Println("Guard:", g)
}

type object rune

const (
	obstacle object = '#'
	visited  object = 'X'
)

type location struct {
	x, y int
	obj  object
}

type direction rune

const (
	up    direction = '^'
	down  direction = 'v'
	left  direction = '<'
	right direction = '>'
)

func (d direction) String() string {
	return string(d)
}

// Turn returns the direction that the guard should turn to
func (d direction) Turn() direction {
	var res direction
	switch d {
	case up:
		res = right
	case right:
		res = down
	case down:
		res = left
	case left:
		res = up
	default:
		return d
	}
	d = res
	return res
}

type guard struct {
	x, y int
	dir  direction // '^><v'
}
