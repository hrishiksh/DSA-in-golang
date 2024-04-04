package algorithms

import (
	"errors"
	"math"
)

func BinarySearch(l []int, t int) (int, error) {
	var startIndex float64 = 0
	var endIndex float64 = float64(len(l) - 1)

	for startIndex <= endIndex {
		midpoint := math.Floor((startIndex + endIndex) / 2)
		// fmt.Printf("Start index: %v, End index: %v, Midpoint: %v, Value: %d, Target: %d \n", startIndex, endIndex, midpoint, l[int(midpoint)], t)

		switch true {
		case l[int(midpoint)] == t:
			return int(midpoint), nil
		case l[int(midpoint)] < t:
			startIndex = midpoint + 1
		case l[int(midpoint)] > t:
			endIndex = midpoint - 1
		}
	}

	return 0, errors.New("target not found")
}
