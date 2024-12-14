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
		log.Fatal("Failed to read input file")
	}
	fileContents := strings.TrimFunc(string(file), unicode.IsSpace)
	stones := newStones(fileContents)

	stones = stones.blinkN(25)
	fmt.Println(stones.count())
	stones = stones.blinkN(50)
	fmt.Println(stones.count())
}

type stones map[int]int

func newStones(s string) stones {
	retVal := make(map[int]int)
	for _, str := range strings.Split(s, " ") {
		num, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			log.Fatal("Couldn't parse int from", str)
		}
		retVal[int(num)] += 1
	}
	return retVal
}

func (s stones) blinkN(n int) stones {
	for i := 0; i < n; i++ {
		s = s.blink()
	}
	return s
}

func (s stones) blink() stones {
	newStones := make(map[int]int)
	for stone, count := range s {
		length := countDigits(stone)
		if stone == 0 {
			newStones[1] += count
		} else if length%2 == 0 {
			l, r := splitStone(stone, length)
			newStones[l] += count
			newStones[r] += count
		} else {
			newStones[stone*2024] += count
		}
	}
	return newStones
}

func (s stones) count() (sum int) {
	for _, count := range s {
		sum += count
	}
	return
}

func countDigits(n int) int {
	workingNum := n
	count := 0
	for workingNum != 0 {
		workingNum = workingNum / 10
		count += 1
	}
	return count
}

func splitStone(n int, length int) (left, right int) {
	splitLen := length / 2
	for i := 0; i < splitLen; i++ {
		digit := (n / pow10(i)) % 10
		right += digit * pow10(i)
	}
	for i := splitLen; i < length; i++ {
		digit := (n / pow10(i)) % 10
		left += digit * pow10(i-splitLen)
	}
	return
}

func pow10(n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num = num * 10
	}
	return num
}
