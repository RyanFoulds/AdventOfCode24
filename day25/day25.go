package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day25/pkg/locks"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file:", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	puzzle := locks.NewPuzzle(fileContents)

	fmt.Println(puzzle.SolvePartOne())
}
