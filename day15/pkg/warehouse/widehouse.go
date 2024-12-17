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
	case '^':
		wh.moveUp()
	case '>':
		wh.moveRight()
	case 'v':
		wh.moveDown()
	case '<':
		wh.moveLeft()
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

func (wh *Widehouse) moveDown() {
	g, i, j := wh.grid, wh.botI, wh.botJ
	cache := map[Coord]struct{}{Coord{i, j}: struct{}{}}
	stack := []Coord{Coord{i, j}}
	var c Coord

	for len(stack) > 0 {
		c, stack = stack[0], stack[1:]
		switch g[c.i][c.j] {
		case '@':
			d := Coord{c.i + 1, c.j}
			cache[d] = struct{}{}
			stack = append(stack, d)
		case '.':
			// Nothing to do.
			continue
		case '#':
			// Can't move so just short-circuit.
			return
		case '[':
			d, r := Coord{c.i + 1, c.j}, Coord{c.i, c.j + 1}
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
			d, l := Coord{c.i + 1, c.j}, Coord{c.i, c.j - 1}
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

	// Move in reverse order of 'i'
	toMove := groupByI(cache)
	keys := make([]int, 0)
	for k, _ := range toMove {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	//	if len(keys) > 0 {
	keys = keys[1:]
	//	}
	for _, k := range keys {
		for _, c := range toMove[k] {
			if wh.grid[c.i][c.j] != '.' {
				wh.grid[c.i+1][c.j] = wh.grid[c.i][c.j]
				wh.grid[c.i][c.j] = '.'
			}
		}
	}
	wh.botI += 1
}

func (wh *Widehouse) moveUp() {
	g, i, j := wh.grid, wh.botI, wh.botJ
	cache := map[Coord]struct{}{Coord{i, j}: struct{}{}}
	stack := []Coord{Coord{i, j}}
	var c Coord

	for len(stack) > 0 {
		c, stack = stack[0], stack[1:]
		switch g[c.i][c.j] {
		case '@':
			u := Coord{c.i - 1, c.j}
			cache[u] = struct{}{}
			stack = append(stack, u)
		case '.':
			// Nothing to do.
			continue
		case '#':
			// Can't move so just short circuit.
			return
		case '[':
			u, r := Coord{c.i - 1, c.j}, Coord{c.i, c.j + 1}
			_, okU := cache[u]
			_, okr := cache[r]
			if !okU {
				cache[u] = struct{}{}
				stack = append(stack, u)
			}
			if !okr {
				cache[r] = struct{}{}
				stack = append(stack, r)
			}
		case ']':
			u, l := Coord{c.i - 1, c.j}, Coord{c.i, c.j - 1}
			_, okU := cache[u]
			_, okl := cache[l]
			if !okU {
				cache[u] = struct{}{}
				stack = append(stack, u)
			}
			if !okl {
				cache[l] = struct{}{}
				stack = append(stack, l)
			}
		}
	}

	// Move in order of 'i'
	toMove := groupByI(cache)
	keys := make([]int, 0)
	for k, _ := range toMove {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//	if len(keys) > 0 {
	keys = keys[1:]
	//	}

	for _, k := range keys {
		for _, c := range toMove[k] {
			if wh.grid[c.i][c.j] != '.' {
				wh.grid[c.i-1][c.j] = wh.grid[c.i][c.j]
				wh.grid[c.i][c.j] = '.'
			}
		}
	}
	wh.botI -= 1
}

func (wh *Widehouse) moveRight() {
	finalCol := wh.botJ + 1
	for isBox(wh.grid[wh.botI][finalCol]) {
		finalCol += 1
	}

	if wh.grid[wh.botI][finalCol] == '#' {
		// Can't move
		return
	}

	for j := finalCol; j > wh.botJ; j-- {
		wh.grid[wh.botI][j] = wh.grid[wh.botI][j-1]
	}
	wh.grid[wh.botI][wh.botJ] = '.'
	wh.botJ += 1
}

func (wh *Widehouse) moveLeft() {
	finalCol := wh.botJ - 1
	for isBox(wh.grid[wh.botI][finalCol]) {
		finalCol -= 1
	}

	if wh.grid[wh.botI][finalCol] == '#' {
		// Can't move
		return
	}

	for j := finalCol; j < wh.botJ; j++ {
		wh.grid[wh.botI][j] = wh.grid[wh.botI][j+1]
	}
	wh.grid[wh.botI][wh.botJ] = '.'
	wh.botJ -= 1
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
