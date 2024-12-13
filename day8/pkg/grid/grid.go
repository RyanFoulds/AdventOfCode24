package grid

import (
	"strings"

	"github.com/RyanFoulds/AdventOfCode24/day8/pkg/fraction"
)

type Coordinate struct {
	i, j int
}

func (c Coordinate) inBounds(h, w int) bool {
	return c.i >= 0 && c.j >= 0 && c.i < h && c.j < w
}

func (c Coordinate) add(other Coordinate) Coordinate {
	return Coordinate{c.i + other.i, c.j + other.j}
}

type Grid struct {
	firstOrderAntinodes map[rune]map[Coordinate]struct{}
	allAntinodes        map[rune]map[Coordinate]struct{}
	antennae            map[rune][]Coordinate
}

func (g Grid) CountAntinodes(firstOrderOnly bool) int {
	uniqueCoords := make(map[Coordinate]struct{})
	var nodes map[rune]map[Coordinate]struct{}
	if firstOrderOnly {
		nodes = g.firstOrderAntinodes
	} else {
		nodes = g.allAntinodes
	}

	for _, coords := range nodes {
		for coord, _ := range coords {
			uniqueCoords[coord] = struct{}{}
		}
	}
	return len(uniqueCoords)
}

func NewGrid(input string) Grid {
	antennae := make(map[rune][]Coordinate)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char == '.' {
				continue
			}
			existing, present := antennae[char]
			if !present {
				existing = make([]Coordinate, 0)
			}
			existing = append(existing, Coordinate{i, j})
			antennae[char] = existing
		}
	}

	h := len(lines)
	w := len(lines[0])

	allAntinodes := make(map[rune]map[Coordinate]struct{})
	firstOrderAntinodes := make(map[rune]map[Coordinate]struct{})
	for frequency, coords := range antennae {
		allAntinodes[frequency] = findAntinodes(frequency, coords, h, w)
		firstOrderAntinodes[frequency] = findFirstOrderAntinodes(frequency, coords, h, w)
	}

	return Grid{firstOrderAntinodes, allAntinodes, antennae}
}

func findFirstOrderAntinodes(frequency rune, antennae []Coordinate, height, width int) map[Coordinate]struct{} {
	retVal := make(map[Coordinate]struct{})

	antennaeCount := len(antennae)
	for i := 0; i < antennaeCount-1; i++ {
		for j := i + 1; j < antennaeCount; j++ {
			first, second := findFirstAntinodes(antennae[i], antennae[j])

			if first.inBounds(height, width) {
				retVal[first] = struct{}{}
			}
			if second.inBounds(height, width) {
				retVal[second] = struct{}{}
			}
		}
	}
	return retVal
}

func findAntinodes(frequency rune, antennae []Coordinate, height, width int) map[Coordinate]struct{} {
	retVal := make(map[Coordinate]struct{})

	antennaeCount := len(antennae)
	for i := 0; i < antennaeCount-1; i++ {
		for j := i + 1; j < antennaeCount; j++ {
			newAntinodes := findAllAntinodes(antennae[i], antennae[j], height, width)
			for k, v := range newAntinodes {
				retVal[k] = v
			}
			//			first, second := findAntinodePair(antennae[i], antennae[j])

			//			if first.inBounds(height, width) {
			//				retVal[first] = struct{}{}
			//			}
			//			if second.inBounds(height, width) {
			//				retVal[second] = struct{}{}
			//			}
		}
	}
	return retVal
}

func findAllAntinodes(first Coordinate, second Coordinate, height, width int) map[Coordinate]struct{} {
	retVal := make(map[Coordinate]struct{})

	di := second.i - first.i
	dj := second.j - first.j
	normDi, normDj := fraction.ReduceFraction(di, dj)
	vector := Coordinate{normDi, normDj}
	reverseVector := Coordinate{-normDi, -normDj}

	coord := first
	for coord.inBounds(height, width) {
		retVal[coord] = struct{}{}
		coord = coord.add(vector)
	}
	coord = first.add(reverseVector)
	for coord.inBounds(height, width) {
		retVal[coord] = struct{}{}
		coord = coord.add(reverseVector)
	}

	return retVal
}

func findFirstAntinodes(first, second Coordinate) (firstResult, secondResult Coordinate) {
	firstResult = Coordinate{2*second.i - first.i, 2*second.j - first.j}
	secondResult = Coordinate{2*first.i - second.i, 2*first.j - second.j}
	return
}
