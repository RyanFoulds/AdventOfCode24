package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var safeCount int
	var kindaSafeCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levelsAsStrings := strings.Fields(scanner.Text())

		report := make([]int, len(levelsAsStrings))
		for index, level := range levelsAsStrings {
			int64Level, err := strconv.ParseInt(level, 10, 0)
			if err != nil {
				os.Exit(1)
			}
			report[index] = int(int64Level)
		}

		unsafeIndex := getFirstUnsafeIndex(report)
		if unsafeIndex < 0 {
			safeCount += 1
			kindaSafeCount += 1
		} else {

			// There are only ever three candidates for removal: the element that triggered the failure, the element before that, of the 0th element.
			fixed1 := make([]int, len(report)-1)
			copy(fixed1, report[:unsafeIndex])
			copy(fixed1[unsafeIndex:], report[unsafeIndex+1:])

			fixed2 := make([]int, len(report)-1)
			copy(fixed2, report[:unsafeIndex-1])
			copy(fixed2[unsafeIndex-1:], report[unsafeIndex:])

			fixed3 := report[1:]

			if isSafe(fixed1) || isSafe(fixed2) || isSafe(fixed3) {
				kindaSafeCount += 1
			}
		}
	}

	fmt.Println(safeCount)
	fmt.Println(kindaSafeCount)
}

func isSafe(report []int) bool {
	return getFirstUnsafeIndex(report) < 0
}

func getFirstUnsafeIndex(report []int) int {
	var priorNegativeDiff bool
	var priorPositiveDiff bool

	last := report[0]
	for i := 1; i < len(report); i++ {
		current := report[i]
		diff := current - last
		last = current

		if diff > 3 || diff < -3 || diff == 0 || (diff > 0 && priorNegativeDiff) || (diff < 0 && priorPositiveDiff) {
			return i
		}

		if diff > 0 {
			priorPositiveDiff = true
		} else {
			// Don't need to consider the diff == 0 case because that would have already triggered returning false above.
			priorNegativeDiff = true
		}

	}

	return -1
}
