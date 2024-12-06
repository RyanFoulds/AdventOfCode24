package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var firstLocations sort.IntSlice
	var secondLocations sort.IntSlice
	left := make(map[int]int)
	right := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		locations := strings.Fields(scanner.Text())

		firstLocationId, err1 := strconv.ParseInt(locations[0], 10, 0)
		secondLocationId, err2 := strconv.ParseInt(locations[1], 10, 0)
		if err1 != nil || err2 != nil {
			fmt.Println(locations)
			os.Exit(1)
		}
		firstLocations = append(firstLocations, int(firstLocationId))
		secondLocations = append(secondLocations, int(secondLocationId))

		// Part 2.
		left[int(firstLocationId)]++
		right[int(secondLocationId)]++
	}

	firstLocations.Sort()
	secondLocations.Sort()
	totalDiff := 0
	for i := 0; i < len(firstLocations); i++ {
		totalDiff += AbsDiff(firstLocations[i], secondLocations[i])
	}
	fmt.Println(totalDiff)

	totalSimilarity := 0
	for key := range left {
		totalSimilarity += key * left[key] * right[key]
	}
	fmt.Println(totalSimilarity)
}

func AbsDiff(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}
