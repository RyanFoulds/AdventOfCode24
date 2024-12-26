package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day24/pkg/logic"
)


func main() {
  filePath := os.Args[1]
  file, err := os.ReadFile(filePath)
  if err != nil {
    log.Fatal("Could not read file:", filePath)
  }
  fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

  puzzle := logic.NewPuzzle(fileContents)
  fmt.Println(puzzle.SolvePartOne())
}

