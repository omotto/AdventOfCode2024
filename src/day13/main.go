package main

import (
	"fmt"
	"math"
	"path/filepath"

	"advent2024/pkg/file"
)

const maxInteger = 9223372036854775807

type coord struct {
	x, y int
}

type Game struct {
	buttonA coord
	buttonB coord
	price   coord
}

func parseInput(s []string) []Game {
	var aX, aY, bX, bY, pX, pY int
	games := make([]Game, (len(s)+1)/4)
	for idx := 0; idx < len(s); idx += 4 {
		_, _ = fmt.Sscanf(s[idx+0], "Button A: X+%d, Y+%d", &aX, &aY)
		_, _ = fmt.Sscanf(s[idx+1], "Button B: X+%d, Y+%d", &bX, &bY)
		_, _ = fmt.Sscanf(s[idx+2], "Prize: X=%d, Y=%d", &pX, &pY)
		games[idx/4] = Game{
			buttonA: coord{x: aX, y: aY},
			buttonB: coord{x: bX, y: bY},
			price:   coord{x: pX, y: pY},
		}
	}
	return games
}

func getPrice(game Game) int {
	minimum := maxInteger
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			x := game.buttonA.x*a + game.buttonB.x*b
			y := game.buttonA.y*a + game.buttonB.y*b
			if x == game.price.x && y == game.price.y && (a*3+b) < minimum {
				minimum = a*3 + b
			}
		}
	}
	if minimum != maxInteger {
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
	pX := game.price.x + delta
	pY := game.price.y + delta
	aX := game.buttonA.x
	aY := game.buttonA.y
	bX := game.buttonB.x
	bY := game.buttonB.y

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
