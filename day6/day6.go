package main

import (
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day6/grid"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	lab := grid.NewGrid(fileContents)
	lab.Walk()
}
