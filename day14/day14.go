package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not load file", filePath)
	}

	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	robotSpecs := strings.Split(fileContents, "\n")
	robots := make([]Robot, len(robotSpecs))
	for i, s := range robotSpecs {
		robots[i] = NewRobot(s)
	}

	safetyFactor := partOne(robots)
	fmt.Println(safetyFactor)

	partTwo(robots)
}

const (
	topLeft = iota
	topRight
	bottomRight
	bottomLeft
	none
)

type Robot struct {
	x, y   int
	vx, vy int
}

func NewRobot(s string) Robot {
	parts := strings.Split(s, " v=")
	positions := strings.Split(parts[0], ",")
	velocities := strings.Split(parts[1], ",")
	x, errX := strconv.ParseInt(positions[0][2:], 10, 0)
	y, errY := strconv.ParseInt(positions[1], 10, 0)
	vx, errVx := strconv.ParseInt(velocities[0], 10, 0)
	vy, errVy := strconv.ParseInt(velocities[1], 10, 0)
	if errX != nil || errY != nil || errVx != nil || errVy != nil {
		log.Fatal("Couldn't parse number!!")
	}

	return Robot{int(x), int(y), int(vx), int(vy)}
}

/*
	Move for n seconds in an X*Y room, return the final normalised coordinates.
*/
func (r Robot) move(n, X, Y int) (int, int) {
	movedX := r.x + n*r.vx
	movedY := r.y + n*r.vy

	normX := movedX % X
	normY := movedY % Y

	if normX < 0 {
		normX = X + normX
	}
	if normY < 0 {
		normY = Y + normY
	}

	return normX, normY
}

func (r Robot) moveN(n, X, Y int) Robot {
	movedX := r.x + n*r.vx
	movedY := r.y + n*r.vy

	normX := movedX % X
	normY := movedY % Y

	if normX < 0 {
		normX = X + normX
	}
	if normY < 0 {
		normY = Y + normY
	}

	return Robot{normX, normY, r.vx, r.vy}
}

func moveManyN(robots []Robot, n, X, Y int) []Robot {
	for i, r := range robots {
		robots[i] = r.moveN(n, X, Y)
	}
	return robots
}

func (r Robot) quadrantAfterMoving(n, X, Y int) int {
	x, y := r.move(n, X, Y)
	if x < X/2 {
		if y < Y/2 {
			return topLeft
		} else if y > Y/2 {
			return bottomLeft
		} else {
			return none
		}
	} else if x > X/2 {
		if y < Y/2 {
			return topRight
		} else if y > Y/2 {
			return bottomRight
		} else {
			return none
		}
	} else {
		return none
	}
}

func partOne(robots []Robot) int {
	robotCounts := make(map[int]int)
	for _, r := range robots {
		quadrant := r.quadrantAfterMoving(100, 101, 103)
		if quadrant != none {
			robotCounts[quadrant] += 1
		}
	}

	product := 1
	for _, count := range robotCounts {
		product *= count
	}
	return product
}

func partTwo(robots []Robot) {
	displayBots := make([]Robot, len(robots))
	copy(displayBots, robots)

	tXoffset := 0
	tYoffset := 0

	// Figure out the step number after which the vertical/horizontal grouping recurr.
	for t := 1; t < 104; t++ {
		displayBots = moveManyN(displayBots, 1, 101, 103)

		mad_x := mad(displayBots, func(r Robot) int { return r.x })
		mad_y := mad(displayBots, func(r Robot) int { return r.y })
		if mad_x < 0.35 {
			tXoffset = t
		}
		if mad_y < 0.35 {
			tYoffset = t
		}
		if tXoffset != 0 && tYoffset != 0 {
			break
		}
	}

	// Find the earliset timestep that sits on both vertical and horizontal groupings at the same time.
	// That's probably our pattern, print it out to show the user, and output the time.
	for i := 0; true; i++ {
		T_c := tXoffset + i*101
		if (T_c-tYoffset)%103 == 0 {
			newBots := moveManyN(robots, T_c, 101, 103)
			displayRobots(newBots, 101, 103)
			fmt.Println(T_c)
			break
		}
	}

}

func mad(robots []Robot, extractor func(r Robot) int) float64 {
	avg := mean(robots, extractor)
	sum := 0.
	for _, r := range robots {
		sum += math.Abs(float64(extractor(r)) - avg)
	}

	return sum / (float64(len(robots)) * avg)
}
func mean(robots []Robot, extractor func(r Robot) int) float64 {
	sum := 0
	for _, r := range robots {
		sum += extractor(r)
	}
	return float64(sum) / float64(len(robots))
}

func displayRobots(robots []Robot, X, Y int) {
	grid := make([][]rune, Y)
	for i := 0; i < Y; i++ {
		grid[i] = make([]rune, X)
		for j := 0; j < X; j++ {
			grid[i][j] = ' '
		}
	}

	for _, r := range robots {
		grid[r.y][r.x] = '#'
	}

	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Print("\n")
	}
}
