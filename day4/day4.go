package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}

	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	lines := strings.Split(fileContents, "\n")

	var wordSearch grid = make([][]rune, len(lines))
	for i := range wordSearch {
		wordSearch[i] = make([]rune, len(lines[i]))
	}

	for i, line := range lines {
		for j, char := range line {
			wordSearch[i][j] = char
		}
	}

	partOne := wordSearch.countSolutions("XMAS")
	partTwo := wordSearch.countXmasBoxes()
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

type grid [][]rune

func (wordSearch grid) countSolutions(term string) int {
	width := len(wordSearch)
	height := len(wordSearch[0])
	var totalCount int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			totalCount += wordSearch.countSolutionsFrom(term, i, j)
		}
	}

	return totalCount
}

func (wordSearch grid) countSolutionsFrom(term string, i, j int) int {
	var matchCount int
	termLength := len(term)
	gridWidth := len(wordSearch[0])
	gridHeight := len(wordSearch)

	// RIGHT
	if j <= gridWidth-termLength {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i][j+k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// DOWN_RIGHT
	if j <= gridWidth-termLength && i <= gridHeight-termLength {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i+k][j+k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// DOWN
	if i <= gridHeight-termLength {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i+k][j] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// DOWN_LEFT
	if j >= termLength-1 && i <= gridHeight-termLength {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i+k][j-k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// LEFT
	if j >= termLength-1 {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i][j-k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// UP_LEFT
	if j >= termLength-1 && i >= termLength-1 {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i-k][j-k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// UP
	if i >= termLength-1 {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i-k][j] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}
	// UP_RIGHT
	if j <= gridWidth-termLength && i >= termLength-1 {
		var match bool = true
		for k, char := range term {
			if char != wordSearch[i-k][j+k] {
				match = false
				break
			}
		}
		if match {
			matchCount += 1
		}
	}

	return matchCount
}

func (wordSearch grid) countXmasBoxes() int {
	searchWidth := len(wordSearch) - 2
	searchHeight := len(wordSearch[0]) - 2
	var totalCount int

	for i := 0; i < searchHeight; i++ {
		for j := 0; j < searchWidth; j++ {
			if wordSearch.isXmasBox(i, j) {
				totalCount += 1
			}
		}
	}
	return totalCount
}

//i, j are the indices at the top-left of the box.
func (g grid) isXmasBox(i, j int) bool {
	if g[i+1][j+1] != 'A' {
		return false
	}

	if g[i][j] == 'S' && g[i+2][j+2] == 'M' {
		if g[i+2][j] == 'M' && g[i][j+2] == 'S' {
			return true
		}
		if g[i+2][j] == 'S' && g[i][j+2] == 'M' {
			return true
		}
	}

	if g[i][j] == 'M' && g[i+2][j+2] == 'S' {
		if g[i+2][j] == 'M' && g[i][j+2] == 'S' {
			return true
		}
		if g[i+2][j] == 'S' && g[i][j+2] == 'M' {
			return true
		}
	}
	return false

	//	return ((g[i][j] == 'S' && g[i+2][j+2] == 'M') ||
	//		(g[i][j] == 'M' && g[i+2][j+2] == 'S')) &&
	//		((g[i+2][j] == 'S' && g[i][j+2] == 'M') ||
	//			(g[i+2][j] == 'M' && g[i][j+2] == 'S'))
}
