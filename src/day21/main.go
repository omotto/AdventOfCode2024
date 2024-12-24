package main

import (
	"advent2024/pkg/file"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

/*
from collections import Counter

codes = open("day21.txt").read().split("\n")

keyp = {c: (i % 3, i // 3) for i, c in enumerate("789456123 0A")}
dirp = {c: (i % 3, i // 3) for i, c in enumerate(" ^A<v>")}


def steps(G: dict[complex, str], s: str, i=1):
    px, py = G["A"]
    bx, by = G[" "]
    res = Counter()
    for c in s:
        npx, npy = G[c]
        f = npx == bx and py == by or npy == by and px == bx
        res[(npx - px, npy - py, f)] += i
        px, py = npx, npy
    return res


def go(n):
    r = 0
    for code in codes:
        res = steps(keyp, code)
        for _ in range(n + 1):
            res = sum((steps(dirp, ("<" * -x + "v" * y + "^" * -y + ">" * x)[:: -1 if f else 1] + "A", res[(x, y, f)]) for x, y, f in res), Counter())
        r += res.total() * int(code[:3])
    return r


print(go(2))
print(go(25))
*/

func getKeyPadCoords(key string) (int, int) {
	switch key {
	case "7":
		return 0, 0
	case "8":
		return 0, 1
	case "9":
		return 0, 2
	case "4":
		return 1, 0
	case "5":
		return 1, 1
	case "6":
		return 1, 2
	case "1":
		return 2, 0
	case "2":
		return 2, 1
	case "3":
		return 2, 2
	case " ":
		return 3, 0
	case "0":
		return 3, 1
	case "A":
		return 3, 2
	}
	return -1, -1
}

func getDirCoords(key string) (int, int) {
	switch key {
	case " ":
		return 0, 0
	case "^":
		return 0, 1
	case "A":
		return 0, 2
	case "<":
		return 1, 0
	case "v":
		return 1, 1
	case ">":
		return 1, 2
	}
	return -1, -1
}

type triple struct {
	dx, dy int
	dir    bool
}

func getSteps(isPad bool, chain string, i int) map[triple]int {
	var okY, okX, badY, badX int
	result := make(map[triple]int)
	if isPad {
		okY, okX = getKeyPadCoords("A")
		badY, badX = getKeyPadCoords(" ")
	} else {
		okY, okX = getDirCoords("A")
		badY, badX = getDirCoords(" ")
	}
	for _, v := range chain {
		var nextY, nextX int
		if isPad {
			nextY, nextX = getKeyPadCoords(string(v))
		} else {
			nextY, nextX = getDirCoords(string(v))
		}
		result[triple{
			dx:  nextX - okX,
			dy:  nextY - okY,
			dir: (nextX == badX && okY == badY) || (nextY == badY && okX == badX),
		}] += i
		okX, okY = nextX, nextY
	}
	return result
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func getSumComplexities(s []string, times int) int {
	result := 0
	for _, code := range s {
		firstSteps := getSteps(true, code, 1) // first KeyPad steps
		for i := 0; i < times+1; i++ {
			newValues := make(map[triple]int)
			for k, v := range firstSteps {
				var left, right, up, down string
				if k.dx < 0 {
					left = strings.Repeat("<", -k.dx)
				} else {
					right = strings.Repeat(">", k.dx)
				}
				if k.dy < 0 {
					up = strings.Repeat("^", -k.dy)
				} else {
					down = strings.Repeat("v", k.dy)
				}
				instructions := left + down + up + right
				if k.dir {
					instructions = Reverse(instructions)
				}
				steps := getSteps(false, instructions+"A", v) // DirectionalPad Steps
				for kk, vv := range steps {
					newValues[kk] += vv
				}
			}
			firstSteps = newValues
		}
		counter := 0
		for _, v := range firstSteps {
			counter += v
		}
		numCode, _ := strconv.Atoi(code[:3])
		result += counter * numCode
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day21/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumComplexities(output, 2))
	fmt.Println(getSumComplexities(output, 25))
}
