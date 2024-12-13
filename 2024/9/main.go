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
	fs = defragBlocks(fs)
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

func mergeAdjBlocks(fs []Block) []Block {
	for i := 0; i < len(fs)-1; i++ {
		if fs[i].id == fs[i+1].id {
			fs[i].size += fs[i+1].size
			fs = slices.Delete(fs, i+1, i+2)
		}
	}
	for fs[len(fs)-1].id == -1 { // if the last blocks are space blocks, remove them
		fs = slices.Delete(fs, len(fs)-1, len(fs))
	}
	return fs
}

func defragBlocks(fs []Block) []Block {
	for id := fs[len(fs)-1].id; id >= 0; id-- {
		mergeAdjBlocks(fs)
		i := slices.IndexFunc(fs, func(b Block) bool { // find the first block with the given id
			return b.id == id
		})
		spc_idx := slices.IndexFunc(fs, func(b Block) bool { // find the first space block
			return b.id == -1 && fs[i].size <= b.size
		})
		if spc_idx > i || spc_idx == -1 { // skip if space block after the current block
			continue
		}
		g := fs[spc_idx].size         // the size of the space block
		fs[spc_idx].id = fs[i].id     // move the block, to fill the space
		fs[spc_idx].size = fs[i].size // move the block, to fill the space
		fs[i].id = -1                 // mark the current block as a space block
		if g > fs[i].size {           // if there's still space left in the space block
			fs = slices.Insert(fs, spc_idx+1, Block{-1, g - fs[i].size}) // insert a new space block to keep remaining space
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
			fs = slices.Delete(fs, i, i+1) // remove the last block, which is now moved
		}
	}
	return fs
}

func checksum(fs []Block) int {
	var sum, pos int
	for _, b := range fs {
		if b.id == -1 { // skip spaces
			pos += b.size
			continue
		}
		for i := 0; i < b.size; i++ {
			sum += b.id * pos
			pos++
		}
	}
	return sum
}
