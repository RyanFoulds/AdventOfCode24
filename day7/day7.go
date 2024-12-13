package main

import (
	"errors"
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
		os.Exit(1)
	}

	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	equationStrings := strings.Split(fileContents, "\n")
	equations := make([]equation, len(equationStrings))

	for i, eq := range equationStrings {
		equations[i] = newEquation(eq)
	}

	partOneSum := 0
	partTwoSum := 0
	for _, eq := range equations {
		if eq.couldBeTrue(false) {
			partOneSum += eq.target
		}
		if eq.couldBeTrue(true) {
			partTwoSum += eq.target
		}
	}
	fmt.Println(partOneSum)
	fmt.Println(partTwoSum)
}

type equation struct {
	target   int
	operands []int
}

func newEquation(line string) equation {
	trimmedLine := strings.TrimFunc(line, unicode.IsSpace)
	parts := strings.Split(trimmedLine, ": ")
	target, err1 := strconv.ParseInt(parts[0], 10, 0)
	if err1 != nil {
		log.Fatalln("Couldn't parse", parts[0], "as an integer")
	}
	ops := strings.Split(parts[1], " ")
	operands := make([]int, len(ops))

	for i, op := range ops {
		parsedOp, err2 := strconv.ParseInt(op, 10, 0)
		if err2 != nil {
			log.Fatalln("Couldn't parse", op, "as an integer")
		}
		operands[i] = int(parsedOp)
	}
	return equation{int(target), operands}
}

func (eq equation) couldBeTrue(allowConcat bool) bool {
	queue := make([]equation, 0)
	queue = append(queue, eq)

	for len(queue) > 0 {
		workingEq := queue[0]
		queue = queue[1:]

		numOfOperands := len(workingEq.operands)
		if numOfOperands > 1 {
			added, errAdd := workingEq.add()
			multiplied, errMul := workingEq.multiply()
			concatted, errCon := workingEq.concat()
			if errAdd != nil || errMul != nil || errCon != nil {
				os.Exit(1)
			}
			if added.operands[0] <= added.target {
				queue = append(queue, added)
			}
			if multiplied.operands[0] <= multiplied.target {
				queue = append(queue, multiplied)
			}
			if allowConcat && concatted.operands[0] <= concatted.target {
				queue = append(queue, concatted)
			}
		} else if numOfOperands == 1 && workingEq.target == workingEq.operands[0] {
			return true
		}
	}

	return false
}

func (eq equation) multiply() (equation, error) {
	numOfIncomingOperands := len(eq.operands)
	if numOfIncomingOperands < 2 {
		return eq, errors.New("Not enough operands to multiply!")
	}
	newOperands := make([]int, numOfIncomingOperands-1)
	if numOfIncomingOperands > 2 {
		copy(newOperands[1:], eq.operands[2:])
	}
	newOperands[0] = eq.operands[0] * eq.operands[1]

	return equation{eq.target, newOperands}, nil
}

func (eq equation) add() (equation, error) {
	numOfIncomingOperands := len(eq.operands)
	if numOfIncomingOperands < 2 {
		return eq, errors.New("Not enough operands to add!")
	}
	newOperands := make([]int, numOfIncomingOperands-1)
	if numOfIncomingOperands > 2 {
		copy(newOperands[1:], eq.operands[2:])
	}
	newOperands[0] = eq.operands[0] + eq.operands[1]

	return equation{eq.target, newOperands}, nil
}

func (eq equation) concat() (equation, error) {
	numOfIncomingOperands := len(eq.operands)
	if numOfIncomingOperands < 2 {
		return eq, errors.New("Not enough operands to concat!")
	}
	newOperands := make([]int, numOfIncomingOperands-1)
	if numOfIncomingOperands > 2 {
		copy(newOperands[1:], eq.operands[2:])
	}
	newOperands[0] = (concatFactor(eq.operands[1]) * eq.operands[0]) + eq.operands[1]

	return equation{eq.target, newOperands}, nil
}

func concatFactor(input int) int {
	val := input
	digitCount := 0
	for val != 0 {
		val = val / 10
		digitCount += 1
	}
	return powOfTen(digitCount)
}

func powOfTen(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	y := powOfTen(n / 2)
	if n%2 == 0 {
		return y * y
	}
	return 10 * y * y
}
