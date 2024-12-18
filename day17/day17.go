package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day17/pkg/cpu"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	computer := cpu.NewComputer(fileContents)
	output := computer.Run()
	fmt.Println(output)

	computer = cpu.NewComputer(fileContents)
	minA := computer.Search()
	fmt.Println(minA)
}
