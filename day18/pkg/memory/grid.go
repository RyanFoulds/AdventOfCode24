package memory

import (
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type BlockedCoords map[Coord]struct{}

func blockedCoords(coords []Coord) BlockedCoords {
	bc := make(BlockedCoords)
	for _, c := range coords {
		bc[c] = struct{}{}
	}
	return bc
}

func AllBlockedCoords(s string) []Coord {
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

func ShortestPath(coords []Coord, start, end Coord) (int, bool) {
	b := blockedCoords(coords)
	visited := map[Coord]int{start: 0}
	q := []Coord{start}
	var node Coord

	for len(q) > 0 {
		node, q = q[0], q[1:]

		for _, c := range node.nextCoords() {
			_, alreadyReached := visited[c]
			if c.isValidStep(b) && !alreadyReached {
				q = append(q, c)
				visited[c] = visited[node] + 1
			}
		}
	}

	distance, validPath := visited[end]
	return distance, validPath
}

func SearchForBlockage(allC []Coord, start, end Coord) Coord {
	l := 1024
	r := len(allC) - 1
	m := (l + r) / 2

	for l != m && r != m {
		_, ok := ShortestPath(allC[:m], start, end)

		if ok {
			l, r, m = m, r, (m+r)/2
		} else {
			l, r, m = l, m, (l+m)/2
		}
	}
	return allC[m]
}
