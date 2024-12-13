package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

const MAX_INT = 9223372036854775807

type Game struct {
	buttonA [2]int
	buttonB [2]int
	price   [2]int
}

func parseLine(s string, idx int) (x int, y int) {
	partsA := strings.Split(s[idx:], ", ")
	x, _ = strconv.Atoi(partsA[0][2:])
	y, _ = strconv.Atoi(partsA[1][2:])
	return x, y
}

func parseInput(s []string) []Game {
	games := make([]Game, (len(s)+1)/4)
	for idx := 0; idx < len(s); idx += 4 {
		aX, aY := parseLine(s[idx], 10)
		bX, bY := parseLine(s[idx+1], 10)
		pX, pY := parseLine(s[idx+2], 7)
		games[idx/4] = Game{
			buttonA: [2]int{aX, aY},
			buttonB: [2]int{bX, bY},
			price:   [2]int{pX, pY},
		}
	}
	return games
}

func getPrice(game Game) int {
	minimum := MAX_INT
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			x := game.buttonA[0]*a + game.buttonB[0]*b
			y := game.buttonA[1]*a + game.buttonB[1]*b
			if x == game.price[0] && y == game.price[1] && (a*3+b) < minimum {
				minimum = a*3 + b
			}
		}
	}
	if minimum != MAX_INT {
		return minimum
	}
	return 0
}

func getTokens(s []string) int {
	games := parseInput(s)
	result := 0
	for _, game := range games {
		result += getPrice(game)
	}
	return result
}

func getPriceByEquation(game Game, delta int) int {
	/*
		2x2 System Linear Equations
		a * aX + b * bX = pX
		a * aY + b * bY = pY
	*/
	pX := game.price[0] + delta
	pY := game.price[1] + delta
	aX := game.buttonA[0]
	aY := game.buttonA[1]
	bX := game.buttonB[0]
	bY := game.buttonB[1]

	a := float64(pX*bY-pY*bX) / float64(aX*bY-aY*bX)
	b := float64(pY*aX-pX*aY) / float64(aX*bY-aY*bX)

	// if there is no decimals is valid
	if a == math.Trunc(a) && b == math.Trunc(b) {
		return int(a*3 + b)
	}
	return 0
}

func getTokens2(s []string, delta int) int {
	games := parseInput(s)
	result := 0
	for _, game := range games {
		result += getPriceByEquation(game, delta)
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day13/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getTokens(output))
	fmt.Println(getTokens2(output, 10000000000000))
}
