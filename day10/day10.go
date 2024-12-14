package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read input file")
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	grid := parseGrid(fileContents)
	score, rating := grid.countScoreAndRating()
	fmt.Println(score)
	fmt.Println(rating)
}

type Coordinate struct {
	i, j int
}

type Grid struct {
	g             [][]int
	height, width int
	cache         map[Coordinate]map[Coordinate]struct{}
	ratingCache   map[Coordinate]int
}

func parseGrid(s string) Grid {
	g := make([][]int, 0)
	cache := make(map[Coordinate]map[Coordinate]struct{})

	for i, line := range strings.Split(s, "\n") {
		g = append(g, make([]int, len(line)))
		for j, char := range line {
			g[i][j] = int(char) - int('0')
		}
	}
	return Grid{g, len(g), len(g[0]), cache, make(map[Coordinate]int)}
}

func (g *Grid) countScoreAndRating() (score, rating int) {

	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if g.g[i][j] == 0 {
				set, ijRating := g.getSetAndRatingFor(Coordinate{i, j})
				score += len(set)
				rating += ijRating
			}
		}
	}
	return
}

func (g *Grid) getSetAndRatingFor(c Coordinate) (map[Coordinate]struct{}, int) {

	set, cacheHit := g.cache[c]
	rating, ratingCacheHit := g.ratingCache[c]
	if cacheHit && ratingCacheHit {
		return set, rating
	}

	newRating := 0
	newCacheEntry := make(map[Coordinate]struct{})
	if g.g[c.i][c.j] == 9 {
		newRating = 1
		newCacheEntry = map[Coordinate]struct{}{c: struct{}{}}
	} else {
		for _, t := range []Coordinate{Coordinate{c.i, c.j - 1}, Coordinate{c.i - 1, c.j},
			Coordinate{c.i, c.j + 1}, Coordinate{c.i + 1, c.j}} {

			if t.i < 0 || t.i > g.height-1 || t.j < 0 || t.j > g.width-1 {
				continue
			}
			if g.g[t.i][t.j]-g.g[c.i][c.j] != 1 {
				continue
			}
			testSet, testRating := g.getSetAndRatingFor(t)
			addAll(newCacheEntry, testSet)
			newRating += testRating
		}
	}

	g.cache[c] = newCacheEntry
	g.ratingCache[c] = newRating
	return newCacheEntry, newRating
}

func addAll(dest, src map[Coordinate]struct{}) {
	for k, v := range src {
		dest[k] = v
	}
}
