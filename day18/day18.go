package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day18/pkg/memory"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	start, end := memory.Coord{0, 0}, memory.Coord{70, 70}
	block := memory.CreateBlockedCoords(fileContents, 1024)
	path, _ := block.ShortestPath(start, end)
	fmt.Println(path)

	blockage := memory.SearchForBlockage(fileContents, start, end)
	fmt.Printf("%d,%d\n", blockage.X, blockage.Y)
}
