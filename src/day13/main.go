package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("src/day13/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getPrices(output))
}
