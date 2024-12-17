package warehouse

import (
	"strings"
)

type Warehouse struct {
	grid       [][]rune
	botI, botJ int
	moves      []rune
}

type Coord struct {
	i, j int
}

func NewWarehouse(s string) Warehouse {
	parts := strings.Split(s, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	moves := strings.Replace(parts[1], "\n", "", -1)
	var botI, botJ int

	grid := make([][]rune, len(gridLines))
	for i, row := range gridLines {
		grid[i] = make([]rune, len(row))
		for j, char := range row {
			grid[i][j] = char
			if char == '@' {
				botI, botJ = i, j
			}
		}
	}

	return Warehouse{grid, botI, botJ, []rune(moves)}
}

func (wh *Warehouse) DoAllTheMoves() {
	for len(wh.moves) > 0 {
		wh.doAMove()
	}
}

func (wh *Warehouse) doAMove() {
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

func (wh *Warehouse) moveUp() {
	finalRow := wh.botI - 1
	for wh.grid[finalRow][wh.botJ] == 'O' {
		finalRow -= 1
	}

	if wh.grid[finalRow][wh.botJ] == '#' {
		// Can't move
		return
	}

	wh.grid[wh.botI][wh.botJ] = '.'
	if finalRow != wh.botI-1 {
		wh.grid[finalRow][wh.botJ] = 'O'
	}
	// Move the bot.
	wh.botI -= 1
	wh.grid[wh.botI][wh.botJ] = '@'
}

func (wh *Warehouse) moveRight() {
	finalCol := wh.botJ + 1
	for wh.grid[wh.botI][finalCol] == 'O' {
		finalCol += 1
	}

	if wh.grid[wh.botI][finalCol] == '#' {
		// Can't move
		return
	}

	wh.grid[wh.botI][wh.botJ] = '.'
	if finalCol != wh.botJ+1 {
		wh.grid[wh.botI][finalCol] = 'O'
	}
	wh.botJ += 1
	wh.grid[wh.botI][wh.botJ] = '@'
}

func (wh *Warehouse) moveDown() {
	finalRow := wh.botI + 1
	for wh.grid[finalRow][wh.botJ] == 'O' {
		finalRow += 1
	}

	if wh.grid[finalRow][wh.botJ] == '#' {
		// Can't move
		return
	}

	wh.grid[wh.botI][wh.botJ] = '.'
	if finalRow != wh.botI+1 {
		wh.grid[finalRow][wh.botJ] = 'O'
	}
	// Move the bot.
	wh.botI += 1
	wh.grid[wh.botI][wh.botJ] = '@'
}

func (wh *Warehouse) moveLeft() {
	finalCol := wh.botJ - 1
	for wh.grid[wh.botI][finalCol] == 'O' {
		finalCol -= 1
	}

	if wh.grid[wh.botI][finalCol] == '#' {
		// Can't move
		return
	}

	wh.grid[wh.botI][wh.botJ] = '.'
	if finalCol != wh.botJ-1 {
		wh.grid[wh.botI][finalCol] = 'O'
	}
	wh.botJ -= 1
	wh.grid[wh.botI][wh.botJ] = '@'
}

func (wh Warehouse) GetSumOfBoxCoords() (sum int) {
	for i, row := range wh.grid {
		for j, char := range row {
			if char == 'O' {
				sum += 100*(i) + j
			}
		}
	}
	return
}
