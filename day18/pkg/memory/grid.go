package memory

import (
	"container/heap"
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type BlockedCoords map[Coord]struct{}

func CreateBlockedCoords(s string, n int) BlockedCoords {
	allC := allCoords(s)
	return blockedCoords(allC[:n])
}

func blockedCoords(coords []Coord) BlockedCoords {
	bc := make(BlockedCoords)
	for _, c := range coords {
		bc[c] = struct{}{}
	}
	return bc
}

func allCoords(s string) []Coord {
	coordStrs := strings.Split(s, "\n")
	allC := make([]Coord, len(coordStrs))

	for i, c := range coordStrs {
		parts := strings.Split(c, ",")
		x, errX := strconv.ParseInt(parts[0], 10, 0)
		y, errY := strconv.ParseInt(parts[1], 10, 0)
		if errX != nil || errY != nil {
			log.Fatal("Couldn't parse int.")
		}
		allC[i] = Coord{int(x), int(y)}
	}

	return allC
}

func (c Coord) isValidStep(b BlockedCoords) bool {
	_, blocked := b[c]
	return !(blocked || c.X < 0 || c.Y < 0 || c.X > 70 || c.Y > 70)
}

func (c Coord) nextCoords() []Coord {
	return []Coord{Coord{c.X + 1, c.Y}, Coord{c.X - 1, c.Y}, Coord{c.X, c.Y + 1}, Coord{c.X, c.Y - 1}}
}

func (b BlockedCoords) ShortestPath(start, end Coord) (int, bool) {
	visited := make(map[Coord]int)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	item := &Item{
		value:    start,
		priority: 0,
	}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Item)

		for _, c := range node.value.nextCoords() {
			_, alreadyReached := visited[c]
			if c.isValidStep(b) && !alreadyReached {
				nextItem := &Item{
					value:    c,
					priority: node.priority - 1,
				}
				heap.Push(&pq, nextItem)
				visited[c] = -1 * nextItem.priority
			}
		}
	}
	distance, validPath := visited[end]
	return distance, validPath
}

func SearchForBlockage(s string, start, end Coord) Coord {
	allC := allCoords(s)
	l := 1024
	r := len(allC) - 1
	m := (l + r) / 2

	for l < r {
		_, ok1 := blockedCoords(allC[:m-1]).ShortestPath(start, end)
		_, ok2 := blockedCoords(allC[:m]).ShortestPath(start, end)
		changed := ok1 != ok2

		if changed {
			return allC[m-1]
		} else if ok1 {
			l, r, m = m, r, (m+r)/2
		} else {
			l, r, m = l, m, (l+m)/2
		}

	}
	return Coord{-1, -1}
}
