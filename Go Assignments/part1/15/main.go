package main

import (
	"log"
)

var number int
var channelEven = make(chan bool)
var channelOdd = make(chan bool)
var channelDone = make(chan bool)

func updateVarWithEvenValue() {
	for i := 0; i <= 10; i += 2 {
		log.Printf("Enter even...")
		<-channelEven
		number = i
		log.Printf("[%d]: %d", i, number)
		if i == 10 {
			channelDone <- true
		} else {
			channelOdd <- true
		}
	}

}

func updateVarWithOddValue() {
	for i := 1; i <= 10; i += 2 {
		log.Printf("Enter odd...")
		<-channelOdd
		number = i
		log.Printf("[%d]: %d", i, number)
		channelEven <- true
	}
}

func main() {
	go updateVarWithOddValue()
	go updateVarWithEvenValue()
	channelEven <- true
	<-channelDone
}
