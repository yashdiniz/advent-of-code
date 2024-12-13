package main

import (
	"aoc2024/internal"
	"log"
	"slices"
	"strconv"
)

func main() {
	lines := internal.OpenInputFile("input.txt").ReadLines()
	dense := lines[0]
	fs := toBlocks(dense)
	fs = compactBlocks(fs)
	sum := checksum(fs)
	log.Println("Checksum:", sum)
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

func compactBlocks(fs []Block) []Block {
	for {
		i := len(fs) - 1
		if fs[i].id == -1 { // skip spaces and remove them
			fs = slices.Delete(fs, i, i+1)
			continue
		}
		// find the first space block
		spc_idx := slices.IndexFunc(fs, func(b Block) bool {
			return b.id == -1
		})
		if spc_idx == -1 { // no space blocks, so we're done compacting
			break
		}

		// g := Block{fs[i].id, fs[i].size - fs[spc_idx].size}
		// if g.size > 0 {
		// 	fs = slices.Replace(fs, spc_idx, spc_idx+1, g)
		// }
		// TODO: fix this
		if fs[spc_idx].size <= fs[i].size { // if the space block is the same size or smaller than the current block
			fs[spc_idx].id = fs[i].id                  // move the block, to fill the space
			fs[i].size = fs[i].size - fs[spc_idx].size // reduce the current block size
			if fs[i].size == 0 {                       // if the current block is now empty, remove it
				fs = slices.Delete(fs, i, i+1)
			}
		} else { // if the space block is larger than the current block
			fs = slices.Replace(fs, spc_idx, spc_idx+1,
				Block{fs[i].id, fs[i].size},              // fill the space with the current block
				Block{-1, fs[spc_idx].size - fs[i].size}, // and keep the remaining space
			)
			fs = slices.Delete(fs, len(fs)-1, len(fs)) // remove the last block, which is now moved
		}
	}
	return fs
}

func checksum(fs []Block) int {
	var sum, pos int
	for _, b := range fs {
		for i := 0; i < b.size; i++ {
			sum += b.id * pos
			pos++
		}
	}
	return sum
}
