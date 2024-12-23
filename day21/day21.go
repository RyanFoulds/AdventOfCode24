package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/RyanFoulds/AdventOfCode24/day21/pkg/robot"
)

func main() {
  filePath := os.Args[1]
  file, err := os.ReadFile(filePath)
  if err != nil {
    log.Fatal("Couldn't load file", filePath)
  }
  
  fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
  codes := strings.Split(fileContents, "\n")

  fmt.Println(robot.Solve(codes, 2))
  fmt.Println(robot.Solve(codes, 25))
}

