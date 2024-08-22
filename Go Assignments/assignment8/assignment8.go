package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cities := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

	var f *os.File
	var err error

	f, err = os.Open("cities")
	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist! Creating file...")
		f, err = os.Create("cities")
		check(err)
	}
	w := bufio.NewWriter(f)

	for i := 0; i < len(cities); i++ {
		_, err := w.WriteString(cities[i] + "\n")
		check(err)
		//fmt.Printf("wrote %d bytes\n", bytes)
	}
	w.Flush()

	sc := bufio.NewScanner(f)
	var city string
	var citiesFromFile []string
	for sc.Scan() {
		city = sc.Text()
		citiesFromFile = append(citiesFromFile, city)
	}

	slices.Sort(citiesFromFile)

	fmt.Println("List of cities from file in alphabetical order:", citiesFromFile)
}
