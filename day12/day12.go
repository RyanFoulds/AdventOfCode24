package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day12/pkg/farm"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read input file")
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	f := farm.NewGarden(fileContents)
	fmt.Println(f.GetCost())
	fmt.Println(f.GetDiscountedCost())
}
