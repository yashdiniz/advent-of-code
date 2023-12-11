package internal

import (
	"bufio"
	"log"
	"os"
)

type InputFile struct {
	file *os.File
}

func OpenInputFile(path string) InputFile {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed at internal.OpenInputFile", err)
	}
	return InputFile{file}
}

func (f InputFile) ReadLines() []string {
	scanner := bufio.NewScanner(f.file)
	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("failed at internal/InputFile.ReadLines", err)
	}
	return out
}
