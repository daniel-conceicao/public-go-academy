package main

import (
	"fmt"
	"slices"
	"sync"
)

func showSliceInfo(prefix string, slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s Slice: %v\n", prefix, slice)
	oddSlice := make([]int, 0)
	evenSlice := make([]int, 0)
	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		} else {
			oddSlice = append(oddSlice, num)
		}
	}
	fmt.Printf("%s Even numbers slice [%d]: %v\n", prefix, len(evenSlice), evenSlice)
	fmt.Printf("%s Odd numbers slice [%d]: %v\n", prefix, len(oddSlice), oddSlice)
}

func main() {
	const count = 10
	array := make([]int, count)
	var wg sync.WaitGroup

	//array initializaion
	for i := range array {
		array[i] = i + 1
	}

	fmt.Println("Original array:", array)

	slice := array[0:count]
	fmt.Println("Original slice:", slice)

	//Sort ascending
	slices.Sort(array[0:count])
	wg.Add(1)
	go showSliceInfo("Ascending", slice, &wg)

	//Sort descending
	revSlice := make([]int, len(slice))
	copy(revSlice, slice)
	slices.Reverse(revSlice)
	wg.Add(1)
	go showSliceInfo("Descending", revSlice, &wg)

	wg.Wait()
}
