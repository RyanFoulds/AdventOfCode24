package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day22/pkg/rand"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	seedStrs := strings.Split(fileContents, "\n")
	seeds := make([]int, len(seedStrs))
	for i, seed := range seedStrs {
		s, err := strconv.ParseInt(seed, 10, 0)
		if err != nil {
			log.Fatal("Couldn't parse int:", seed)
		}
		seeds[i] = int(s)
	}

	fmt.Println(solvePartOne(seeds))
  fmt.Println(solvePartTwo(seeds))
}

func solvePartOne(seeds []int) (sum int) {
	for _, n := range seeds {
		sum += rand.NextN(n, 2000)
	}
	return
}

func solvePartTwo(seeds []int) (max int) {
	sequencePayoffs := rand.MonkeyFromSeed(seeds, 2000)
	for _, v := range sequencePayoffs {
		if v > max {
			max = v
		}
	}
	return
}
