package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}

	var sum int
	scanner := bufio.NewScanner(file)
	var fullFile string
	for scanner.Scan() {
		line := scanner.Text()
		fullFile += line
		sum += parse(line)
	}
	fmt.Println(sum)
	fmt.Println(parseConditionally(fullFile))
}

func parse(line string) int {
	var sumOfProducts int

	candidates := strings.Split(line, "mul(")
	for _, candidate := range candidates {
		closingIndex := strings.Index(candidate, ")")
		if closingIndex < 3 {
			continue
		}
		tuple := candidate[:closingIndex]

		args := strings.Split(tuple, ",")
		if len(args) != 2 {
			continue
		}

		left, errLeft := strconv.ParseInt(args[0], 10, 0)
		right, errRight := strconv.ParseInt(args[1], 10, 0)
		if errLeft != nil || errRight != nil {
			continue
		}
		product := int(left * right)
		sumOfProducts += product
	}
	return sumOfProducts
}

func parseConditionally(line string) int {
	var sum int
	splits := strings.Split(line, "do()")
	for _, split := range splits {
		closingIndex := strings.Index(split, "don't()")
		if closingIndex > 0 {
			split = split[:closingIndex]
		}
		sum += parse(split)
	}
	return sum
}
