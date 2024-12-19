package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day19/pkg/pattern"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	puzzle := pattern.NewPuzzle(fileContents)

	fmt.Println(puzzle.SolvePartOne())
	fmt.Println(puzzle.SolvePartTwo())
}
