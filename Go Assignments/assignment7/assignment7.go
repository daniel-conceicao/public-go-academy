package main

import (
	"fmt"
	"math/rand/v2"
)

func rollDice() int {
	return rand.IntN(6+1-1) + 1
}
func printRollInfo(rollCount, dice1, dice2 int, outcome string) {
	fmt.Printf("Roll[%d] [%d, %d]: %s\n", rollCount, dice1, dice2, outcome)
}

func main() {
	numRolls := 50

	for i := 0; i < numRolls; i++ {
		dice1Roll := rollDice()
		dice2Roll := rollDice()
		switch sumRolls := dice1Roll + dice2Roll; sumRolls {
		case 7, 11:
			printRollInfo(i, dice1Roll, dice2Roll, "NATURAL")
		case 2:
			printRollInfo(i, dice1Roll, dice2Roll, "SNAKE-EYES-CRAPS")
		case 3, 12:
			printRollInfo(i, dice1Roll, dice2Roll, "LOSS CRAPS")
		default:
			printRollInfo(i, dice1Roll, dice2Roll, "NEUTRAL")
		}
	}
}
