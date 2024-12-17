package maze

import (
	//	"fmt"
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
	n       node
	score   int
	nodeLog []node
}

type node struct {
	i, j, di, dj int
}

type coord struct {
	i, j int
}

func (d deer) rotateCW() deer {
	newNode := node{d.n.i, d.n.j, d.n.dj, -d.n.di}

	newLog := make([]node, len(d.nodeLog)+1)
	copy(newLog, d.nodeLog)
	newLog[len(d.nodeLog)] = newNode

	return deer{newNode, d.score + 1000, newLog}
}

func (d deer) rotateACW() deer {
	newNode := node{d.n.i, d.n.j, -d.n.dj, d.n.di}

	newLog := make([]node, len(d.nodeLog)+1)
	copy(newLog, d.nodeLog)
	newLog[len(d.nodeLog)] = newNode

	return deer{newNode, d.score + 1000, newLog}
}

func (d deer) step() deer {
	newI, newJ := d.n.i+d.n.di, d.n.j+d.n.dj
	newNode := node{newI, newJ, d.n.di, d.n.dj}

	newLog := make([]node, len(d.nodeLog)+1)
	copy(newLog, d.nodeLog)
	newLog[len(d.nodeLog)] = newNode

	return deer{newNode, d.score + 1, newLog}
}

func (m *Maze) Search() int {
	startingNode := node{m.si, m.sj, 0, 1}
	startingDeer := deer{startingNode, 0, []node{startingNode}}
	minScore := map[node]int{startingNode: 0}
	winningDeer := make(map[node][]deer)
	stack := []deer{startingDeer}
	var d deer

	for len(stack) > 0 {
		d, stack = stack[0], stack[1:]

		a := d.step()

		if m.grid[a.n.i][a.n.j] != '#' {
			currentMin, aOk := minScore[a.n]
			if !aOk || currentMin > a.score {
				stack = append(stack, a)
				minScore[a.n] = a.score
				winningDeer[a.n] = []deer{a}
			} else if currentMin == a.score {
				stack = append(stack, a)
				winningDeer[a.n] = append(winningDeer[a.n], a)
			}
		}

		b := d.rotateCW()
		bCurrentMin, bOk := minScore[b.n]
		if !bOk || bCurrentMin > b.score {
			stack = append(stack, b)
			minScore[b.n] = b.score
			winningDeer[b.n] = []deer{b}
		} else if bCurrentMin == b.score {
			stack = append(stack, b)
			winningDeer[b.n] = append(winningDeer[b.n], b)
		}

		c := d.rotateACW()
		cCurrentMin, cOk := minScore[c.n]
		if !cOk || cCurrentMin > c.score {
			stack = append(stack, c)
			minScore[c.n] = c.score
			winningDeer[c.n] = []deer{c}
		} else if cCurrentMin == c.score {
			stack = append(stack, c)
			winningDeer[c.n] = append(winningDeer[c.n], c)
		}
	}

	// Should have exhausted the maze by now??
	// Four possible ending nodes, facing:
	up, right, down, left := node{m.ei, m.ej, -1, 0}, node{m.ei, m.ej, 0, 1}, node{m.ei, m.ej, 1, 0}, node{m.ei, m.ej, 0, -1}
	minUp, minRight, minDown, minLeft := minScore[up], minScore[right], minScore[down], minScore[left]

	min := min(minUp, minRight, minDown, minLeft)

	allWinningDeer := make([]deer, 0)
	for _, n := range []node{up, right, down, left} {
		if minScore[n] == min {
			allWinningDeer = append(allWinningDeer, winningDeer[n]...)
		}
	}
	m.winningDeer = allWinningDeer

	return min
}

func (m Maze) CountTilesOnWinningPaths() int {
	tileSet := make(map[coord]struct{})
	for _, d := range m.winningDeer {
		for _, n := range d.nodeLog {
			tileSet[coord{n.i, n.j}] = struct{}{}
		}
	}
	return len(tileSet)
}
