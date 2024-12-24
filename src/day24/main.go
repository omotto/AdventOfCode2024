package main

import (
	"advent2024/pkg/file"
	"bytes"
	"fmt"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func execute(keyValues map[string]int, operand1, operand2, operation string) int {
	a, b := keyValues[operand1], keyValues[operand2]
	switch operation {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	}
	return -1
}

func getDecimalNum(s []string) int {
	var (
		key         string
		value, next int
	)
	keyValue := make(map[string]int)
	for idx, line := range s {
		if len(line) > 2 {
			parts := strings.Split(line, ": ")
			value, _ = strconv.Atoi(parts[1])
			key = parts[0]
			keyValue[key] = value
		} else {
			next = idx
			break
		}
	}
	var operand1, operand2, operation, resultOp string
	for {
		repeat := false
		for idx := next + 1; idx < len(s); idx++ {
			_, _ = fmt.Sscanf(s[idx], "%s %s %s -> %s", &operand1, &operation, &operand2, &resultOp)
			_, ok1 := keyValue[resultOp]
			_, ok2 := keyValue[operand1]
			_, ok3 := keyValue[operand2]
			if !ok1 && ok2 && ok3 {
				keyValue[resultOp] = execute(keyValue, operand1, operand2, operation)
				repeat = true
			}
		}
		if !repeat {
			break
		}
	}
	resultSlice := make([]string, 64)
	for k, v := range keyValue {
		if k[0] == 'z' {
			index, _ := strconv.Atoi(k[1:])
			resultSlice[index] = strconv.Itoa(v)
		}
	}
	var resultStr bytes.Buffer
	for idx := len(resultSlice) - 1; idx >= 0; idx-- {
		resultStr.WriteString(resultSlice[idx])
	}
	result, _ := strconv.ParseInt(resultStr.String(), 2, 64)
	return int(result)
}

func getOperation(operand1, operand2, operator string, operations []string) string {
	idx := slices.IndexFunc(operations, func(str string) bool {
		return strings.HasPrefix(str, fmt.Sprintf("%s %s %s", operand1, operator, operand2)) ||
			strings.HasPrefix(str, fmt.Sprintf("%s %s %s", operand2, operator, operand1))
	})
	if idx == -1 {
		return ""
	}
	return strings.Split(operations[idx], " -> ")[1]
}

func getSwappedBits(s []string) string {
	var (
		swapped []string
		carry,
		nextCarry,
		nextRes string
	)
	for bitCount := 0; bitCount < 45; bitCount++ {
		res1 := getOperation(fmt.Sprintf("x%02d", bitCount), fmt.Sprintf("y%02d", bitCount), "XOR", s)
		int1 := getOperation(fmt.Sprintf("x%02d", bitCount), fmt.Sprintf("y%02d", bitCount), "AND", s)
		if bitCount > 0 {
			int2 := getOperation(carry, res1, "AND", s)
			if int2 == "" {
				res1, int1 = int1, res1
				swapped = append(swapped, res1, int1)
				int2 = getOperation(carry, res1, "AND", s)
			}
			nextRes = getOperation(carry, res1, "XOR", s)
			if res1[0] == 'z' {
				res1, nextRes = nextRes, res1
				swapped = append(swapped, res1, nextRes)
			}
			if int1[0] == 'z' {
				int1, nextRes = nextRes, int1
				swapped = append(swapped, int1, nextRes)
			}
			if int2[0] == 'z' {
				int2, nextRes = nextRes, int2
				swapped = append(swapped, int2, nextRes)
			}
			nextCarry = getOperation(int2, int1, "OR", s)
		}
		if bitCount == 0 {
			carry = int1
		} else {
			if nextCarry[0] == 'z' && nextCarry != "z45" {
				nextCarry, nextRes = nextRes, nextCarry
				swapped = append(swapped, nextCarry, nextRes)
			}
			carry = nextCarry
		}
	}
	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

func main() {
	absPathName, _ := filepath.Abs("src/day24/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getDecimalNum(output))
	fmt.Println(getSwappedBits(output))
}
