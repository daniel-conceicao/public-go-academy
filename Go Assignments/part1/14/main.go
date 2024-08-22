package main

import (
	"log"
	"time"
)

var number int

var oddValue int
var evenValue int
var count int

func updateVarAndPrintNextEvenNumber(number int) {
	for i := 0; i < 100; i++ {
		evenValue += 2
		number = evenValue
		count++
		log.Printf("[%d]: %d\n", count+1, number)
	}
}

func updateVarAndPrintNextOddNumber(number int) {
	for i := 0; i < 100; i++ {
		oddValue += 2
		number = oddValue
		count++
		log.Printf("[%d]: %d\n", count, number)
	}
}

func main() {
	count = 0
	oddValue = 1
	evenValue = 0

	go updateVarAndPrintNextEvenNumber(number)
	go updateVarAndPrintNextOddNumber(number)

	time.Sleep(10 * time.Second)

}
