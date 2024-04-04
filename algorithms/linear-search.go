package algorithms

import (
	"errors"
)

// Linear search algorithm takes a sorted list and iterate
// when an element of the list match with the target, it returns
// the index of that element
func LinearSearch(list []int, target int) (int, error) {
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			return i, nil
		}
	}
	return 0, errors.New("target not found")
}
