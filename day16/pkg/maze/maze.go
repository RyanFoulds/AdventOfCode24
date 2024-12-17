package maze

import (
	"strings"
)

type Maze struct {
	grid        [][]rune
	si, sj      int
	ei, ej      int
	winningDeer []deer
}

func NewMaze(s string) Maze {
	gridLines := strings.Split(s, "\n")

	var ei, ej, si, sj int
	grid := make([][]rune, len(gridLines))
	for i, row := range gridLines {
		grid[i] = make([]rune, len(row))
		for j, char := range row {
			grid[i][j] = char
			if char == 'S' {
				si, sj = i, j
			} else if char == 'E' {
				ei, ej = i, j
			}
		}
	}
	return Maze{grid, si, sj, ei, ej, nil}
}

type deer struct {
	n     node
	score int
}

type node struct {
	i, j, di, dj int
}

type coord struct {
	i, j int
}

func (d deer) rotateCW() deer {
	newNode := node{d.n.i, d.n.j, d.n.dj, -d.n.di}
	return deer{newNode, d.score + 1000}
}

func (d deer) rotateACW() deer {
	newNode := node{d.n.i, d.n.j, -d.n.dj, d.n.di}
	return deer{newNode, d.score + 1000}
}

func (d deer) step(reverse bool) deer {
	var newI, newJ int
	if reverse {
		newI, newJ = d.n.i-d.n.di, d.n.j-d.n.dj
	} else {
		newI, newJ = d.n.i+d.n.di, d.n.j+d.n.dj
	}
	newNode := node{newI, newJ, d.n.di, d.n.dj}
	return deer{newNode, d.score + 1}
}

func (m Maze) Search() (int, int) {
	eNode := node{m.si, m.sj, 0, 1}
	sNodes := []node{node{m.ei, m.ej, 0, 1}, node{m.ei, m.ej, 0, -1}, node{m.ei, m.ej, 1, 0}, node{m.ei, m.ej, -1, 0}}

	forwardScores := m.search(false, []node{eNode})
	backwardScores := m.search(true, sNodes)

	winningScore := backwardScores[eNode]

	onPath := make(map[coord]struct{})
	for n, fs := range forwardScores {
		if fs+backwardScores[n] == winningScore {
			onPath[coord{n.i, n.j}] = struct{}{}
		}
	}

	return winningScore, len(onPath)
}

func (m Maze) search(reverse bool, startingNodes []node) map[node]int {
	stack := make([]deer, 0)
	minScore := make(map[node]int)

	for _, node := range startingNodes {
		stack = append(stack, deer{node, 0})
		minScore[node] = 0
	}

	var d deer
	for len(stack) > 0 {
		d, stack = stack[0], stack[1:]

		a := d.step(reverse)
		if m.grid[a.n.i][a.n.j] != '#' {
			currentMin, aOk := minScore[a.n]
			if !aOk || currentMin > a.score {
				stack = append(stack, a)
				minScore[a.n] = a.score
			}
		}

		b := d.rotateCW()
		bCurrentMin, bOk := minScore[b.n]
		if !bOk || bCurrentMin > b.score {
			stack = append(stack, b)
			minScore[b.n] = b.score
		}

		c := d.rotateACW()
		cCurrentMin, cOk := minScore[c.n]
		if !cOk || cCurrentMin > c.score {
			stack = append(stack, c)
			minScore[c.n] = c.score
		}
	}
	return minScore
}
