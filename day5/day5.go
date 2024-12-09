package main

import (
	"fmt"
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

	parts := strings.Split(fileContents, "\n\n")
	updates := strings.Split(parts[1], "\n")
	rules := createRulesMap(strings.Split(parts[0], "\n"))

	var validMiddlePagesSum int
	var invalidMiddlePagesSum int
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			validMiddlePagesSum += extractMiddlePage(strings.Split(update, ","))
		} else {
			sortedUpdate := sort(update, rules)
			invalidMiddlePagesSum += extractMiddlePage(sortedUpdate)
		}
	}
	fmt.Println(validMiddlePagesSum)
	fmt.Println(invalidMiddlePagesSum)
}

func createRulesMap(rules []string) map[string]map[string]struct{} {
	rulesMap := make(map[string]map[string]struct{})

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		key := splitRule[0]
		value := splitRule[1]
		elem, ok := rulesMap[key]

		var newElement map[string]struct{}
		if ok {
			newElement = elem
		} else {
			newElement = make(map[string]struct{})
		}
		newElement[value] = struct{}{}
		rulesMap[key] = newElement
	}
	return rulesMap
}

func isValidUpdate(update string, rules map[string]map[string]struct{}) bool {
	pages := strings.Split(update, ",")
	for i, page := range pages {
		rule, ok := rules[page]
		if ok {
			for _, priorPage := range pages[:i] {
				_, prohibited := rule[priorPage]
				if prohibited {
					return false
				}
			}
		}
	}
	return true
}

func extractMiddlePage(update []string) int {
	index := len(update) / 2
	value, err := strconv.ParseInt(update[index], 10, 0)
	if err != nil {
		os.Exit(1)
	}
	return int(value)
}

func sort(update string, rules map[string]map[string]struct{}) []string {
	// TODO: Implement this shit properly.
	return strings.Split(update, ",")
}
