package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day20/pkg/racetrack"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	puzzle := racetrack.NewPuzzle(fileContents)

	fmt.Println(puzzle.SolvePartOne())
	fmt.Println(puzzle.SolvePartTwo())
}
