package main

import (
	"advent2024/pkg/file"
	"fmt"
	"strconv"
	"strings"

	"path/filepath"
)

func getMapOrderingRules(s []string) (map[int][]int, int) {
	orderListRules := make(map[int][]int)
	for idx, rule := range s {
		if len(rule) < 3 {
			return orderListRules, idx
		}
		parts := strings.Split(rule, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		if _, ok := orderListRules[a]; ok {
			orderListRules[a] = append(orderListRules[a], b)
		} else {
			orderListRules[a] = []int{b}
		}
	}
	return orderListRules, -1
}

func getUpdates(s []string, start int) [][]int {
	updates := make([][]int, len(s)-start)
	for i := start; i < len(s); i++ {
		parts := strings.Split(s[i], ",")
		updates[i-start] = make([]int, len(parts))
		for j, part := range parts {
			updates[i-start][j], _ = strconv.Atoi(part)
		}
	}
	return updates
}

func validate(src, cmp int, rules map[int][]int, lessMore int) bool {
	if lessMore < 0 { // less
		for _, value := range rules[cmp] {
			if value == src {
				return true
			}
		}
	} else { // more
		for _, value := range rules[src] {
			if value == cmp {
				return true
			}
		}
	}
	return false
}

func validateUpdate(update []int, rules map[int][]int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := 0; j < len(update); j++ {
			if i != j {
				if !validate(update[i], update[j], rules, j-i) {
					return false
				}
			}
		}
	}
	return true
}

func getSumMiddlePageNumberFromRightUpdates(s []string) int {
	rules, start := getMapOrderingRules(s)
	updates := getUpdates(s, start+1)
	result := 0
	for _, update := range updates {
		if validateUpdate(update, rules) {
			result += update[len(update)/2]
		}
	}
	return result
}

func fixUpdate(update []int, rules map[int][]int) []int {
	fixedUpdate := make([]int, len(update))
	for i := 0; i < len(update); i++ {
		values := rules[update[i]]
		sum := 0
		for j := 0; j < len(update); j++ {
			if i != j {
				for k := 0; k < len(values); k++ {
					if values[k] == update[j] {
						sum++
						break
					}
				}
			}
		}
		fixedUpdate[len(update)-sum-1] = update[i]
	}
	return fixedUpdate
}

func getSumMiddlePageNumberFromFixedUpdates(s []string) int {
	rules, start := getMapOrderingRules(s)
	updates := getUpdates(s, start+1)
	result := 0
	for _, update := range updates {
		if !validateUpdate(update, rules) {
			fixedUpdate := fixUpdate(update, rules)
			result += fixedUpdate[len(fixedUpdate)/2]
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day05/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumMiddlePageNumberFromRightUpdates(output))
	fmt.Println(getSumMiddlePageNumberFromFixedUpdates(output))
}
