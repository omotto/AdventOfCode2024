package main

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"advent2024/pkg/file"
)

const (
	MULT = '1'
	SUM  = '0'
	CONC = '2'
)

func getResult(values []int, operators string) int {
	var result int
	for x := 0; x < len(operators); x++ {
		if x == 0 {
			result = values[x]
		}
		switch operators[x] {
		case MULT:
			result = result * values[x+1]
		case SUM:
			result = result + values[x+1]
		case CONC:
			result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, values[x+1]))
		}
	}
	return result
}

func isValid(result int, values []int) bool {
	numComb := int(math.Pow(float64(2), float64(len(values)-1)))
	for i := 0; i < numComb; i++ {
		operators := strconv.FormatInt(int64(i), 2)
		ops := strings.Repeat("0", len(values)-len(operators)-1) + operators
		if getResult(values, ops) == result {
			return true
		}
	}
	return false
}

func validate(result int, values []int, numOps, currentDepth int, combination []int) bool {
	for i := 0; i < numOps; i++ {
		newCombination := slices.Clone(combination)
		newCombination = append(newCombination, i)
		if len(values)-1 == currentDepth {
			ops := strings.Trim(strings.Replace(fmt.Sprint(newCombination), " ", "", -1), "[]")
			if getResult(values, ops) == result {
				return true
			}
		} else {
			if validate(result, values, numOps, currentDepth+1, newCombination) {
				return true
			}
		}
	}
	return false
}

func getValidRows1(s []string) int {
	result := 0
	for _, line := range s {
		p1 := strings.Split(line, ": ")
		key, _ := strconv.Atoi(p1[0])
		p2 := strings.Split(p1[1], " ")
		values := make([]int, len(p2))
		for idx, value := range p2 {
			values[idx], _ = strconv.Atoi(value)
		}
		if isValid(key, values) {
			result += key
		}
	}
	return result
}

func getValidRows2(s []string) int {
	var (
		mu     sync.Mutex
		result = 0
		eg     = errgroup.Group{}
		//wg sync.WaitGroup
	)
	eg.SetLimit(200)
	//wg.Add(len(s))
	for _, line := range s {
		//go func() {
		eg.Go(func() error {
			p1 := strings.Split(line, ": ")
			key, _ := strconv.Atoi(p1[0])
			p2 := strings.Split(p1[1], " ")
			values := make([]int, len(p2))
			for idx, value := range p2 {
				values[idx], _ = strconv.Atoi(value)
			}
			numOps := 3 // SUM, MULT, CONC
			if validate(key, values, numOps, 1, []int{}) {
				mu.Lock()
				result += key
				mu.Unlock()
			}
			//defer wg.Done()
			return nil
		})
		//}()
	}
	_ = eg.Wait()
	//wg.Wait()
	return result

}
func main() {
	absPathName, _ := filepath.Abs("src/day07/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getValidRows1(output))
	fmt.Println(getValidRows2(output))
}
