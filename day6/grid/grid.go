package grid

import (
	"fmt"
	"strings"
)

type direction int

const (
	up = iota
	right
	down
	left
)

var next map[direction]direction = map[direction]direction{up: right, right: down, down: left, left: up}
var current map[rune]direction = map[rune]direction{'^': up, '>': right, 'v': down, '<': left}

type Coordinate struct {
	i, j int
}

type Grid struct {
	grid                         [][]rune
	currentLocation              Coordinate
	currentDirection             direction
	visitedLocations             map[Coordinate]map[direction]int
	possibleObstructionLocations map[Coordinate]struct{}
}

func NewGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	var grid [][]rune = make([][]rune, len(lines))
	var currentLocation Coordinate
	var currentDirection direction
	visitedLocations := make(map[Coordinate]map[direction]int)
	possibleObstructionLocations := make(map[Coordinate]struct{})

	for i := range grid {
		grid[i] = make([]rune, len(lines[i]))
	}
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
			direction, isDirection := current[char]
			if isDirection {
				grid[i][j] = '.' // Don't fill in the security guard as a permanent fixture.
				currentDirection = direction
				currentLocation = Coordinate{i, j}
				addToMap(visitedLocations, currentLocation, currentDirection)
			}
		}
	}

	return Grid{grid, currentLocation, currentDirection, visitedLocations, possibleObstructionLocations}
}

func (g *Grid) Walk() {
	savedStartingLoc := Coordinate{g.currentLocation.i, g.currentLocation.j}
	savedStartingDir := g.currentDirection

	g.walk()

	// Part 1.
	fmt.Println(g.CountLocations())

	// Part 2
	for loc, val := range g.visitedLocations {
		for dir, _ := range val {
			candidateLocation := nextLocation(loc, dir)
			if candidateLocation.isOutOfBounds(g) {
				continue
			}
			tempDir := dir
			count := 0
			for g.grid[candidateLocation.i][candidateLocation.j] == '#' {
				tempDir = next[dir]
				candidateLocation = nextLocation(loc, tempDir)
				count += 1
				if count > 3 {
					break
				}
			}

			g.reset(savedStartingLoc, savedStartingDir)
			oldVal := g.grid[candidateLocation.i][candidateLocation.j]
			g.grid[candidateLocation.i][candidateLocation.j] = '#'
			if g.walk() {
				g.possibleObstructionLocations[candidateLocation] = struct{}{}
			}
			g.grid[candidateLocation.i][candidateLocation.j] = oldVal
		}
	}
	fmt.Println(g.CountPossibleObstacleLocations())
}

func (g *Grid) walk() (looped bool) {
	for {
		exited, looped := g.step()
		if exited || looped {
			return looped
		}
	}
}

func (g *Grid) step() (exited, looped bool) {
	candidateLocation := nextLocation(g.currentLocation, g.currentDirection)

	if candidateLocation.isOutOfBounds(g) {
		return true, false
	}
	if g.visitedLocations[candidateLocation][g.currentDirection] > 0 {
		return false, true
	}

	if g.grid[candidateLocation.i][candidateLocation.j] == '#' {
		g.currentDirection = next[g.currentDirection]
		return g.step()
	} else {
		g.currentLocation = candidateLocation
		addToMap(g.visitedLocations, candidateLocation, g.currentDirection)
	}
	return false, false
}

func (c Coordinate) isOutOfBounds(g *Grid) bool {
	return c.i < 0 ||
		c.i >= len(g.grid) ||
		c.j < 0 ||
		c.j >= len(g.grid[0])
}

func (g Grid) CountLocations() int {
	return len(g.visitedLocations)
}

func (g Grid) CountPossibleObstacleLocations() int {
	return len(g.possibleObstructionLocations)
}

func nextLocation(currentLocation Coordinate, d direction) Coordinate {
	switch d {
	case up:
		return Coordinate{currentLocation.i - 1, currentLocation.j}
	case right:
		return Coordinate{currentLocation.i, currentLocation.j + 1}
	case down:
		return Coordinate{currentLocation.i + 1, currentLocation.j}
	case left:
		return Coordinate{currentLocation.i, currentLocation.j - 1}
	}
	return currentLocation
}

func addToMap(locationHistory map[Coordinate]map[direction]int, location Coordinate, d direction) {
	existingValue, present := locationHistory[location]
	if !present {
		existingValue = make(map[direction]int)
	}
	existingValue[d] += 1
	locationHistory[location] = existingValue
}

func (g *Grid) reset(startingLocation Coordinate, startingDirection direction) {
	g.currentLocation = startingLocation
	g.currentDirection = startingDirection
	g.visitedLocations = make(map[Coordinate]map[direction]int)
}
