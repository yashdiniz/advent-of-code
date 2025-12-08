package main

import (
	"aoc2025/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := internal.OpenInputFile("./input.txt")
	lines := input.ReadLines()

	// parse all operators
	ops_at := len(lines) - 1
	var ops []rune
	for _, op := range lines[ops_at] {
		if op == ' ' {
			continue
		}
		ops = append(ops, op)
	}

	var opnds [][]int
	var acc []int
	// transpose the operand lines (each column is the operand)
	for col := 0; col < len(lines[ops_at]); col++ {
		v := ""
		for i := range lines[:ops_at] {
			v += string(lines[i][col])
		}
		if strings.TrimSpace(v) == "" {
			opnds = append(opnds, acc)
			acc = nil
			continue
		}
		opnd, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			log.Panic("failed at parse")
		}
		acc = append(acc, opnd)
	}
	opnds = append(opnds, acc)

	log.Print(opnds)
	sum := 0
	for i, op := range ops {
		log.Print(string(op), opnds[i])
		switch op { // get top of stack
		case '+':
			acc := 0
			for _, v := range opnds[i] {
				acc += v
			}
			sum += acc
		case '*':
			acc := 1
			for _, v := range opnds[i] {
				acc *= v
			}
			sum += acc
		}
		ops = ops[1:] // pop top of stack
	}
	log.Print("Answer: ", sum)
}
