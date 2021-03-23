package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Pair struct {
	firstVal  int
	secondVal int
}

func main() {
	a := []int{-1, 1, -2, 2, -3, 3, -4, -5, -25, 50, 252, 12412, 235235, 35235, -12412}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	fmt.Printf("Testing range of %d numbers \n", len(a))

	pos, neg := DivideNumbers(a)

	start := time.Now()
	FindPairs(pos, neg)
	duration := time.Since(start)
	fmt.Println("Time lapsed: ", duration)
}

func DivideNumbers(nums []int) (pos []int, neg []int) {
	var negNums []int
	var posNums []int
	for _, v := range nums {
		if v < 0 {
			negNums = append(negNums, v)
		} else {
			posNums = append(posNums, v)
		}
	}

	return posNums, negNums
}

func FindPairs(pos []int, neg []int) {

	sort.Ints(pos)
	sort.Ints(neg)

	var pairs []Pair

	if len(neg) > len(pos) {
		for _, val := range pos {
			var pair Pair
			if BinarySearch(neg, -val) != -1 {
				pair.firstVal = val
				pair.secondVal = -val
				pairs = append(pairs, pair)
			}

		}
	}

	if len(pos) > len(neg) {
		for _, val := range neg {
			var pair Pair
			if BinarySearch(pos, -val) != -1 {
				pair.firstVal = val
				pair.secondVal = -val
				pairs = append(pairs, pair)
			}
		}
	}

	fmt.Printf("There were found %d pairs. Pairs: %d \n", len(pairs), pairs)
}

func BinarySearch(array []int, target int) int {
	startIndex := 0
	endIndex := len(array) - 1
	midIndex := len(array) / 2
	for startIndex <= endIndex {

		value := array[midIndex]

		if value == target {
			return midIndex
		}

		if value > target {
			endIndex = midIndex - 1
			midIndex = (startIndex + endIndex) / 2
			continue
		}

		startIndex = midIndex + 1
		midIndex = (startIndex + endIndex) / 2
	}

	return -1
}
