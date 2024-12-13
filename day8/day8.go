package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day8/pkg/grid"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	antennaGrid := grid.NewGrid(fileContents)

	fmt.Println(antennaGrid.CountAntinodes(true))
	fmt.Println(antennaGrid.CountAntinodes(false))
}
