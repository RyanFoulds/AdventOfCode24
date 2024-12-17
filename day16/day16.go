package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day16/pkg/maze"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	m := maze.NewMaze(fileContents)
	lowScore := m.Search()
	fmt.Println(lowScore)
	tileCount := m.CountTilesOnWinningPaths()
	fmt.Println(tileCount)
}
