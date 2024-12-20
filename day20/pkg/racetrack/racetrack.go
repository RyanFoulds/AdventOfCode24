package racetrack

import (
	"container/heap"
	"strings"
)

type cheatRecord struct {
	i1, j1, i2, j2 int
}

type grid [][]rune

func (g grid) outOfRange(i, j int) bool {
	return i < 0 || j < 0 || i >= len(g) || j >= len(g[i])
}

type Puzzle struct {
	g      grid
	si, sj int
	ei, ej int
}

type programState struct {
	i, j int
	c    cheatRecord
}

func (p programState) canCheat() bool {
	return p.c.i1 == -1
}

type coord struct {
	i, j int
}

func nextItemFunc(maxCheatTime int) func(item Item, g grid) []Item {
	return func(item Item, g grid) []Item {
		retval := make([]Item, 0)

		for x := -maxCheatTime; x <= maxCheatTime; x++ {
			for y := abs(x) - maxCheatTime; y <= maxCheatTime-abs(x); y++ {
				manhatten := abs(x) + abs(y)
				if manhatten == 0 {
					continue
				} else if manhatten == 1 {
					if !g.outOfRange(item.value.i+x, item.value.j+y) &&
						g[item.value.i+x][item.value.j+y] != '#' {
						n := Item{
							value:    programState{item.value.i + x, item.value.j + y, item.value.c},
							priority: item.priority + 1,
						}
						retval = append(retval, n)
					}
				} else if item.value.canCheat() {
					newI, newJ := item.value.i+x, item.value.j+y
					if !g.outOfRange(newI, newJ) &&
						g[newI][newJ] != '#' {
						n := Item{
							value:    programState{newI, newJ, cheatRecord{item.value.i, item.value.j, newI, newJ}},
							priority: item.priority + manhatten,
						}
						retval = append(retval, n)
					}
				}
			}
		}
		return retval
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func NewPuzzle(s string) Puzzle {
	var si, sj, ei, ej int
	rows := strings.Split(s, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = make([]rune, len(row))
		for j, char := range row {
			grid[i][j] = char
			if char == 'E' {
				ei, ej = i, j
			} else if char == 'S' {
				si, sj = i, j
			}
		}
	}

	return Puzzle{grid, si, sj, ei, ej}
}

func (puzzle Puzzle) SolvePartOne() int {
	return puzzle.djikstraWithCheats(nextItemFunc(2), 100)
}

func (puzzle Puzzle) SolvePartTwo() int {
	return puzzle.djikstraWithCheats(nextItemFunc(20), 100)
}

func (puzzle Puzzle) djikstraWithCheats(nextSteps func(Item, grid) []Item, targetSaving int) int {
	cheatersCache := puzzle.createCache()
	maxResult := cheatersCache[coord{puzzle.si, puzzle.sj}] - targetSaving

	// Initialise the PriorityQueue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	startingItem := &Item{
		value:    programState{puzzle.si, puzzle.sj, cheatRecord{-1, -1, -1, -1}},
		priority: 0,
	}
	heap.Push(&pq, startingItem)

	seen := map[programState]struct{}{startingItem.value: struct{}{}}
	results := make(map[int]int)

	// Do the search.
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)

		if p.priority > maxResult {
			continue
		}
		if p.value.i == puzzle.ei && p.value.j == puzzle.ej {
			results[p.priority]++
		} else {
			for _, nextItem := range nextSteps(*p, puzzle.g) {
				_, ok := seen[nextItem.value]
				if !ok {
					cachedRemaining, reachable := cheatersCache[coord{nextItem.value.i, nextItem.value.j}]

					if nextItem.value.canCheat() {
						heap.Push(&pq, &nextItem)
						seen[nextItem.value] = struct{}{}
					} else if reachable {
						// A program that has already cheated doesn't need to compute the rest of the path, just add the remainder from the cache.
						results[nextItem.priority+cachedRemaining]++
					}
				}
			}
		}
	}

	// Found all routes now, so count how many are under the threshold.
	// Can't rely on the above search short-circuiting because of the addition of cached distances.
	counter := 0
	for k, v := range results {
		if k <= maxResult {
			counter += v
		}
	}
	return counter
}

// Creates a cache of the time taken to reach the end from every reachable location without cheating.
// Use the result as a cache for determining determining the total path length of routes that have just used their cheat.
func (puzzle Puzzle) createCache() map[coord]int {
	nextSteps := nextItemFunc(1)
	startingItem := Item{
		value:    programState{puzzle.ei, puzzle.ej, cheatRecord{-2, -1, -1, -1}}, // -2 prevents cheating
		priority: 0,
	}
	queue := []Item{startingItem}
	seen := map[coord]int{coord{startingItem.value.i, startingItem.value.j}: 0}

	var item Item
	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]

		for _, nextItem := range nextSteps(item, puzzle.g) {
			c := coord{nextItem.value.i, nextItem.value.j}
			_, ok := seen[c]
			if !ok {
				queue = append(queue, nextItem)
				seen[c] = nextItem.priority
			}
		}
	}
	return seen
}
