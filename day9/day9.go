package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day9/pkg/filesystem"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Problem reading file at", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)

	// Part one
	simpleFs := filesystem.SimpleFromString(fileContents)
	fmt.Println(simpleFs.Checksum())

	// Part two
	fs := filesystem.FromString(fileContents)
	fs.MoveFiles()
	fmt.Println(fs.Checksum())
}
