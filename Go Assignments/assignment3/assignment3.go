package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	const lowerRefNumber = 1
	const higherRefNumber = 10
	var number int
	fmt.Printf("Insert a number between %d and %d: ", lowerRefNumber, higherRefNumber)
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}

	result := "is"

	if number < lowerRefNumber || number > higherRefNumber {
		result = "is not"
	}

	fmt.Printf("%d %s between %v and %v", number, result, lowerRefNumber, higherRefNumber)

}
