package farm

import (
	"strings"
)

type Garden struct {
	regions []region
	grid    [][]rune
}

type region struct {
	plant     rune
	plots     map[coordinate]struct{}
	area      int
	perimeter int
}

type coordinate struct {
	i, j int
}

func NewGarden(s string) Garden {
	lines := strings.Split(s, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	retVal := Garden{make([]region, 0), grid}
	retVal.floodAll()

	return retVal
}

func (g Garden) GetCost() (totalCost int) {
	for _, region := range g.regions {
		regCost := region.area * region.perimeter
		totalCost += regCost
	}
	return
}

func (g Garden) GetDiscountedCost() (cost int) {
	for _, region := range g.regions {
		sides := g.countSidesInRegion(region)
		cost += region.area * sides
	}
	return
}

func (g *Garden) floodAll() {
	for i, line := range g.grid {
		for j, _ := range line {
			g.floodFrom(coordinate{i, j})
		}
	}
}

func (g *Garden) floodFrom(c coordinate) {
	if g.visited(c) {
		return
	}
	plant := g.grid[c.i][c.j]

	plots := map[coordinate]struct{}{c: struct{}{}}
	perimeter := 0

	stack := make([]coordinate, 0)
	stack = append(stack, c)

	for len(stack) > 0 {
		var next coordinate
		next, stack = stack[0], stack[1:]

		candidates := []coordinate{coordinate{next.i, next.j - 1}, coordinate{next.i - 1, next.j},
			coordinate{next.i, next.j + 1}, coordinate{next.i + 1, next.j}}

		for _, candidate := range candidates {
			if g.canFloodTo(candidate, plant) {
				_, ok := plots[candidate]
				if !ok {
					stack = append(stack, candidate)
					plots[candidate] = struct{}{}
				}
			} else {
				perimeter += 1
			}
		}
	}

	reg := region{plant, plots, len(plots), perimeter}
	g.regions = append(g.regions, reg)
}

func (g Garden) countSidesInRegion(r region) int {
	cornerCount := 0
	for c, _ := range r.plots {
		cornerCount += g.cornerCountTopLeft(c)
		cornerCount += g.cornerCountTopRight(c)
		cornerCount += g.cornerCountBottomLeft(c)
		cornerCount += g.cornerCountBottomRight(c)
	}
	return cornerCount
}

func (g Garden) cornerCountTopLeft(c coordinate) int {
	plant := g.grid[c.i][c.j]

	eqRight := g.canFloodTo(coordinate{c.i, c.j + 1}, plant)
	eqDown := g.canFloodTo(coordinate{c.i + 1, c.j}, plant)
	eqDiag := g.canFloodTo(coordinate{c.i + 1, c.j + 1}, plant)
	if (!eqRight && !eqDown) || (eqRight && eqDown && !eqDiag) {
		return 1
	} else {
		return 0
	}
}

func (g Garden) cornerCountTopRight(c coordinate) int {
	plant := g.grid[c.i][c.j]

	eqUp := g.canFloodTo(coordinate{c.i, c.j - 1}, plant)
	eqDown := g.canFloodTo(coordinate{c.i + 1, c.j}, plant)
	eqDiag := g.canFloodTo(coordinate{c.i + 1, c.j - 1}, plant)
	if (!eqUp && !eqDown) || (eqUp && eqDown && !eqDiag) {
		return 1
	} else {
		return 0
	}
}

func (g Garden) cornerCountBottomLeft(c coordinate) int {
	plant := g.grid[c.i][c.j]

	eqRight := g.canFloodTo(coordinate{c.i, c.j + 1}, plant)
	eqUp := g.canFloodTo(coordinate{c.i - 1, c.j}, plant)
	eqDiag := g.canFloodTo(coordinate{c.i - 1, c.j + 1}, plant)
	if (!eqRight && !eqUp) || (eqRight && eqUp && !eqDiag) {
		return 1
	} else {
		return 0
	}
}

func (g Garden) cornerCountBottomRight(c coordinate) int {
	plant := g.grid[c.i][c.j]

	eqLeft := g.canFloodTo(coordinate{c.i, c.j - 1}, plant)
	eqUp := g.canFloodTo(coordinate{c.i - 1, c.j}, plant)
	eqDiag := g.canFloodTo(coordinate{c.i - 1, c.j - 1}, plant)
	if (!eqLeft && !eqUp) || (eqLeft && eqUp && !eqDiag) {
		return 1
	} else {
		return 0
	}
}

func (g Garden) canFloodTo(c coordinate, plant rune) bool {
	return !(c.i < 0 || c.j < 0 ||
		c.i >= len(g.grid) || c.j >= len(g.grid[0]) ||
		g.grid[c.i][c.j] != plant)
}

func (g Garden) visited(c coordinate) bool {
	for _, region := range g.regions {
		_, ok := region.plots[c]
		if ok {
			return true
		}
	}
	return false
}
