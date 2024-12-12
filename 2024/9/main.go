package main

import (
	"aoc2024/internal"
	"log"
	"strconv"
)

func main() {
	lines := internal.OpenInputFile("test.txt").ReadLines()
	dense := lines[0]

	log.Println(toBlocks(dense))
}

type Block struct {
	id   int
	size int
}

func toBlocks(dense string) []Block {
	var fs []Block
	for id := 0; id <= len(dense)/2; id++ {
		i := id * 2
		blk_sz, err := strconv.Atoi(string(dense[i]))
		if err != nil {
			panic(err)
		}
		if id <= len(dense)/2-1 {
			spc_sz, err := strconv.Atoi(string(dense[i+1]))
			if err != nil {
				panic(err)
			}
			fs = append(fs, Block{id, blk_sz})
			if spc_sz > 0 {
				fs = append(fs, Block{-1, spc_sz})
			}
		} else {
			fs = append(fs, Block{id, blk_sz})
		}
	}
	return fs
}
