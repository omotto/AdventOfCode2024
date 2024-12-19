package main

import (
	"advent2024/pkg/file"
	"fmt"
	"golang.org/x/sync/errgroup"
	"path/filepath"
	"strings"
	"sync/atomic"
)

func parseInput(s []string) ([]string, []string) {
	towels := strings.Split(s[0], ", ")
	designs := make([]string, len(s)-2)
	for i := 2; i < len(s); i++ {
		designs[i-2] = s[i]
	}
	return towels, designs
}

func isPossible(design string, towels []string, visited map[string]int) int {
	if v, ok := visited[design]; ok {
		return v
	}
	possible := 0
	for _, towel := range towels {
		// Discard if towel pattern doesn't fit
		if len(towel) <= len(design) {
			// if towel pattern is found at the beginning of towel
			if strings.Index(design, towel) == 0 {
				// if full design belongs to towel pattern
				if len(towel) == len(design) {
					possible++
					continue
				}
				possible += isPossible(design[len(towel):], towels, visited)
			}
		}
	}
	visited[design] = possible
	return possible
}

func getNumDesigns(s []string) int {
	towels, designs := parseInput(s)
	result := 0
	possible := make(map[string]int)
	for _, design := range designs {
		if isPossible(design, towels, possible) > 0 {
			result++
		}
	}
	return result
}

func getNumPossibleDesigns(s []string) int {
	towels, designs := parseInput(s)
	result := int64(0)
	eg := errgroup.Group{}
	eg.SetLimit(100)
	for _, design := range designs {
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(isPossible(design, towels, map[string]int{})))
			return nil
		})
	}
	_ = eg.Wait()
	return int(result)
}

func main() {
	absPathName, _ := filepath.Abs("src/day19/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumDesigns(output))
	fmt.Println(getNumPossibleDesigns(output))
}
