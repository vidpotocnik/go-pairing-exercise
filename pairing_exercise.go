package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pair struct {
	valPos int
	valNeg int
}

func main() {
	a := MakeRange(-423455, 500000)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	fmt.Printf("Testing range of %d numbers \n", len(a))

	start := time.Now()
	FindPairs(a)
	duration := time.Since(start)
	fmt.Println("Time lapsed: ", duration)
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func FindPairs(nums []int) {
	sortedNumbers := MergeSort(nums)
	var pairs []Pair

	for i := 0; i < len(sortedNumbers); i++ {
		if sortedNumbers[i] < 0 {
			if sortedNumbers[i] + (-sortedNumbers[i]) == 0 {
				var pair Pair
				pair.valPos = sortedNumbers[i]
				pair.valNeg = -sortedNumbers[i]
				pairs = append(pairs, pair)
			}
		}
	}
	fmt.Printf("There were found %d pairs. \n", len(pairs))
}

func MergeSort(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}

	mid := len(elements) / 2
	left := elements[:mid]
	right := elements[mid:]

	size, i, j := len(left)+len(right), 0, 0

	slice := make([]int, size, size)

	for el := 0; el < size; el++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[el] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[el] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[el] = left[i]
			i++
		} else {
			slice[el] = right[j]
			j++
		}
	}
	return slice
}
