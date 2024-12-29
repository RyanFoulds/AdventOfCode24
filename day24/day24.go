package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day24/pkg/logic"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not read file:", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	puzzle := logic.NewPuzzle(fileContents)
	g := logic.CreateGraph(puzzle)
	p1 := puzzle.SolvePartOne()
	fmt.Println(p1)

	// I'm lazy, so put this is a directory that's already in the .gitignore, even if it doesn't really make sense.
	outputFile, _ := os.Create("bin/adder-graph.gv")
	_ = draw.DOT(g, outputFile)
}
