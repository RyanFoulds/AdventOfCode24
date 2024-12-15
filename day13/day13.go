package main

import (
	"fmt"
	"log"
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
	machineStrings := strings.Split(fileContents, "\n\n")
	totalCost := 0
	totalGiantCost := 0

	for _, machineString := range machineStrings {
		m := NewMachine(machineString)
		m2 := Machine{m.A, m.B, Button{m.P.x + 10000000000000, m.P.y + 10000000000000}}

		totalCost += m.solve()
		totalGiantCost += m2.solve()
	}

	fmt.Println(totalCost)
	fmt.Println(totalGiantCost)
}

type Button struct {
	x, y int
}

func NewButton(s string) Button {
	components := strings.Split(strings.Split(s, ": ")[1], ", ")
	x, errX := strconv.ParseInt(components[0][2:], 10, 0)
	y, errY := strconv.ParseInt(components[1][2:], 10, 0)
	if errX != nil || errY != nil {
		log.Fatal("Couldn't parse the numbers.")
	}
	return Button{int(x), int(y)}
}

type Machine struct {
	A, B, P Button
}

func NewMachine(s string) Machine {
	lines := strings.Split(s, "\n")
	if len(lines) != 3 {
		log.Fatal("Invalid machine!")
	}
	A := NewButton(lines[0])
	B := NewButton(lines[1])
	P := NewButton(lines[2])
	return Machine{A, B, P}
}

func (m Machine) solve() int {

	// Handle all the funky cases where A, B, P have some similarity.
	if areSimilar(m.P, m.A) {
		if areSimilar(m.A, m.B) {
			// A, B, P are all similar. Find the cheapest combo.
			log.Println("All vectors are similar for machine:", m, "Finding cheapest solution is not yet implemented in this case.")
			return 0
		}
		// P, A are similar but B is not, the general case solution below covers this, co-efficient of B must be 0.
	} else if areSimilar(m.P, m.B) {
		// B, P similar but A is not => 0 or 1 solutions depending on if P is an exact multiple of B, co-efficient of A must be 0.
		if m.P.x%m.B.x == 0 {
			// In the odd case that B.x > P.x there are no solutions, so we'd want to return 0 anyway.
			return m.P.x / m.B.x
		} else {
			return 0
		}
	} else if areSimilar(m.A, m.B) {
		// A,B similar but P is not => no solutions possible.
		return 0
	}

	// Start ordinary solution assuming no similarity.
	bNumerator := m.P.y*m.A.x - m.P.x*m.A.y
	bDenominator := m.B.y*m.A.x - m.B.x*m.A.y

	if bDenominator != 0 && bNumerator%bDenominator == 0 {
		b := bNumerator / bDenominator
		aNumerator := m.P.x - b*m.B.x
		if m.A.x != 0 && aNumerator%m.A.x == 0 {
			a := aNumerator / m.A.x
			return 3*a + b
		}
	}

	// Catch the case where a single solution exists despite A.x being 0
	if m.A.x == 0 && m.P.x%m.B.x == 0 {
		b := m.P.x / m.B.x
		aNumerator := m.P.y - b*m.B.y
		if m.A.y != 0 && aNumerator%m.A.y == 0 {
			a := aNumerator / m.A.y
			return 3*a + b
		}
	}

	return 0
}

func areSimilar(A, B Button) bool {
	return A.x/A.y == B.x/B.y &&
		(A.x%A.y)*B.y == (B.x%B.y)*A.y
}
