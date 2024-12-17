package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day15/pkg/warehouse"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	wh := warehouse.NewWarehouse(fileContents)
	wh.DoAllTheMoves()
	partOne := wh.GetSumOfBoxCoords()
	fmt.Println(partOne)

	wide := warehouse.NewWidehouse(fileContents)
	wide.DoAllTheMoves()
	partTwo := wide.GetSumOfBoxCoords()
	fmt.Println(partTwo)
}
