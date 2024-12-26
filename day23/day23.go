package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day23/pkg/network"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Couldn't read file:", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	links := strings.Split(fileContents, "\n")

	network.ProcessLinks(links)
	network.FindNetworks()
	fmt.Println(network.CountNetworks())
	fmt.Println(network.FindBiggestNetwork())

}
