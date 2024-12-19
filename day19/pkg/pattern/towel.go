package pattern

import (
	"strings"
)

type towels []string

type patterns []string

type Puzzle struct {
	t towels
	p patterns
}

func (puzzle Puzzle) SolvePartOne() (count int) {
	cache := make(map[string]bool)
	for _, pattern := range puzzle.p {
		if isPossible(pattern, puzzle.t, cache) {
			count++
		}
	}
	return
}

func (puzzle Puzzle) SolvePartTwo() (count int) {
	cache := make(map[string]int)
	for _, pattern := range puzzle.p {
		count += waysPossible(pattern, puzzle.t, cache)
	}
	return
}

func NewPuzzle(s string) Puzzle {
	parts := strings.Split(s, "\n\n")
	t := strings.Split(parts[0], ", ")
	p := strings.Split(parts[1], "\n")
	return Puzzle{t, p}
}

func isPossible(pattern string, ts towels, cache map[string]bool) bool {
	b, ok := cache[pattern]
	if ok {
		return b
	}

	for _, t := range ts {
		if t == pattern {
			return true
		} else if strings.HasPrefix(pattern, t) {
			isPoss := isPossible(strings.TrimPrefix(pattern, t), ts, cache)
			if isPoss {
				cache[pattern] = true
				return true
			}
		}
	}
	cache[pattern] = false
	return false
}

func waysPossible(pattern string, ts towels, cache map[string]int) (ways int) {

	w, ok := cache[pattern]
	if ok {
		return w
	}

	for _, t := range ts {
		if t == pattern {
			ways++
		} else if strings.HasPrefix(pattern, t) {
			ways += waysPossible(strings.TrimPrefix(pattern, t), ts, cache)
		}
	}
	cache[pattern] = ways
	return
}
