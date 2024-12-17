package warehouse

import (
	"sort"
	"strings"
)

type Widehouse struct {
	grid       [][]rune
	botI, botJ int
	moves      []rune
}

func NewWidehouse(s string) Widehouse {
	parts := strings.Split(s, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	moves := strings.Replace(parts[1], "\n", "", -1)
	var botI, botJ int

	grid := make([][]rune, len(gridLines))
	for i, row := range gridLines {
		grid[i] = make([]rune, 2*len(row))
		for j, char := range row {
			var first, second rune
			switch char {
			case '#':
				first, second = '#', '#'
			case 'O':
				first, second = '[', ']'
			case '@':
				first, second = '@', '.'
				botI, botJ = i, 2*j
			case '.':
				first, second = '.', '.'
			}
			grid[i][2*j] = first
			grid[i][2*j+1] = second
		}
	}

	return Widehouse{grid, botI, botJ, []rune(moves)}
}

func (wh *Widehouse) DoAllTheMoves() {
	for len(wh.moves) > 0 {
		wh.doAMove()
	}
}

func (wh *Widehouse) doAMove() {
	var move rune
	move, wh.moves = wh.moves[0], wh.moves[1:] // pop a move
	switch move {
	case '^', 'v':
		wh.moveVertical(move == '^')
	case '>', '<':
		wh.moveHorizontal(move == '>')
	}
}

func (wh Widehouse) GetSumOfBoxCoords() (sum int) {
	for i, row := range wh.grid {
		for j, char := range row {
			if char == '[' {
				sum += 100*(i) + j
			}
		}
	}
	return
}

func (wh *Widehouse) moveHorizontal(right bool) {
	var move func(j int) int
	var rev func(j int) int
	if right {
		move = func(j int) int { return j + 1 }
		rev = func(j int) int { return j - 1 }
	} else {
		move = func(j int) int { return j - 1 }
		rev = func(j int) int { return j + 1 }
	}

	finalCol := move(wh.botJ)
	for isBox(wh.grid[wh.botI][finalCol]) {
		finalCol = move(finalCol)
	}

	if wh.grid[wh.botI][finalCol] == '#' {
		// Can't move
		return
	}

	for j := finalCol; (right && j > wh.botJ) || (!right && j < wh.botJ); j = rev(j) {
		wh.grid[wh.botI][j] = wh.grid[wh.botI][rev(j)]
	}

	wh.grid[wh.botI][wh.botJ] = '.'
	wh.botJ = move(wh.botJ)
}

func (wh *Widehouse) moveVertical(up bool) {
	var next func(Coord) Coord
	if up {
		next = func(c Coord) Coord { return Coord{c.i - 1, c.j} }
	} else {

		next = func(c Coord) Coord { return Coord{c.i + 1, c.j} }
	}

	g, i, j := wh.grid, wh.botI, wh.botJ
	cache := map[Coord]struct{}{Coord{i, j}: struct{}{}}
	stack := []Coord{Coord{i, j}}
	var c Coord

	for len(stack) > 0 {
		c, stack = stack[0], stack[1:]
		switch g[c.i][c.j] {
		case '@':
			d := next(c)
			cache[d] = struct{}{}
			stack = append(stack, d)
		case '.':
			// Nothing to do.
			continue
		case '#':
			// Can't move so just short-circuit.
			return
		case '[':
			d, r := next(c), Coord{c.i, c.j + 1}
			_, okD := cache[d]
			_, okr := cache[r]
			if !okD {
				cache[d] = struct{}{}
				stack = append(stack, d)
			}
			if !okr {
				cache[r] = struct{}{}
				stack = append(stack, r)
			}
		case ']':
			d, l := next(c), Coord{c.i, c.j - 1}
			_, okD := cache[d]
			_, okl := cache[l]
			if !okD {
				cache[d] = struct{}{}
				stack = append(stack, d)
			}
			if !okl {
				cache[l] = struct{}{}
				stack = append(stack, l)
			}
		}
	}

	// Important to move in the correct order, robot needs to be last.
	toMove := groupByI(cache)
	keys := make([]int, 0)
	for k, _ := range toMove {
		keys = append(keys, k)
	}
	if up {
		sort.Ints(keys)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	}

	for _, k := range keys {
		for _, c := range toMove[k] {
			t := next(c)
			if wh.grid[c.i][c.j] != '.' {
				wh.grid[t.i][t.j] = wh.grid[c.i][c.j]
				wh.grid[c.i][c.j] = '.'
			}
		}
	}
	wh.botI = next(Coord{i, j}).i
}

func isBox(r rune) bool {
	return r == '[' || r == ']'
}

func groupByI(m map[Coord]struct{}) map[int][]Coord {
	grouped := make(map[int][]Coord)
	for c, _ := range m {
		if grouped[c.i] == nil {
			grouped[c.i] = make([]Coord, 0)
		}
		grouped[c.i] = append(grouped[c.i], c)
	}
	return grouped
}
