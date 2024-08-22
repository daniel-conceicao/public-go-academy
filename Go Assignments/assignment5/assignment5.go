package main

import (
	"fmt"
	"math/rand/v2"
)

type sumNumbers struct {
	singleDigitNumberCount, doubleDigitNumberCount, threeDigitNumberCount int
}

func (sum sumNumbers) getSingleDigitNumbers(slice *[]int) {
	sum.generateRandomNumbers(sum.singleDigitNumberCount, 0, 9, slice)
}

func (sum sumNumbers) getDoubleDigitNumbers(slice *[]int) {
	sum.generateRandomNumbers(sum.doubleDigitNumberCount, 10, 99, slice)
}

func (sum sumNumbers) getThreeDigitNumbers(slice *[]int) {
	sum.generateRandomNumbers(sum.threeDigitNumberCount, 100, 999, slice)
}

func (sum sumNumbers) generateRandomNumbers(count, min, max int, slice *[]int) {
	for i := 0; i < count; i++ {
		randomNum := rand.IntN(max+1-min) + min
		*slice = append(*slice, randomNum)
		fmt.Println("Added random number:", randomNum, *slice)
	}
}

func (sum sumNumbers) getTerms(hasSingleDigit, hasDoubleDigitsNumbers, hasThreeDigitsNumbers bool, slice *[]int) {

	if hasSingleDigit {
		sum.getSingleDigitNumbers(slice)
	}

	if hasDoubleDigitsNumbers {
		sum.getDoubleDigitNumbers(slice)
	}

	if hasThreeDigitsNumbers {
		sum.getThreeDigitNumbers(slice)
	}
}

func main() {
	singleDigitNumberCount := 3
	doubleDigitNumberCount := 3
	threeDigitNumberCount := 3

	terms := make([]int, 0)

	sumNumbers := sumNumbers{singleDigitNumberCount, doubleDigitNumberCount, threeDigitNumberCount}
	sumNumbers.getTerms(true, true, true, &terms)
	fmt.Println(terms)
	sum := 0

	for _, num := range terms {
		sum += num
	}

	fmt.Println("Total sum: ", sum)

}
