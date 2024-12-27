package locks

import (
	"strings"
)

type Puzzle struct {
	keys  []key
	locks []lock
}

type key [5]int

type lock struct {
	pins     [5]int
	capacity int
}

func NewPuzzle(s string) Puzzle {
	locks := make([]lock, 0)
	keys := make([]key, 0)

	parts := strings.Split(s, "\n\n")
	for _, part := range parts {
		grid := createGrid(part)
		capacity := len(grid)
		pins := [5]int{0, 0, 0, 0, 0}
		for _, row := range grid {
			for j, char := range row {
				if char == '#' {
					pins[j]++
				}
			}
		}
		if grid[0][0] == '#' {
			locks = append(locks, lock{pins, capacity})
		} else {
			keys = append(keys, pins)
		}
	}
	return Puzzle{keys, locks}
}

func createGrid(s string) [][]rune {
	grid := make([][]rune, 0)
	rows := strings.Split(s, "\n")
	for i, row := range rows {
		grid = append(grid, make([]rune, len(row)))
		for j, char := range row {
			grid[i][j] = char
		}
	}
	return grid
}

func fits(k key, l lock) bool {
	for i := 0; i < 5; i++ {
		if k[i]+l.pins[i] > l.capacity {
			return false
		}
	}
	return true
}

func (p Puzzle) SolvePartOne() (sum int) {
	for _, key := range p.keys {
		for _, lock := range p.locks {
			if fits(key, lock) {
				sum++
			}
		}
	}
	return
}
