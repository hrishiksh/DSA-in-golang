package main

import (
	"fmt"

	"github.com/hrishiksh/DSA-in-golang/algorithms"
)

func main() {
	var intList []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// i, err := algorithms.LinearSearch(intList, 7)
	i, err := algorithms.BinarySearch(intList, 9)
	if err != nil {
		fmt.Printf("Target not found in the list \n")
		return
	}
	fmt.Printf("Target found at index %d\n", i)

}
