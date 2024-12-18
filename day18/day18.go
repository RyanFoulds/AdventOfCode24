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
	allCoordinates := memory.AllBlockedCoords(fileContents)
	start, end := memory.Coord{0, 0}, memory.Coord{70, 70}

	pathLengthAt1024, _ := memory.ShortestPath(allCoordinates[:1024], start, end)
	blockage := memory.SearchForBlockage(allCoordinates, start, end)

	fmt.Println(pathLengthAt1024)
	fmt.Printf("%d,%d\n", blockage.X, blockage.Y)
}
