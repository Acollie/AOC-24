package main

import (
	"AOC-24/pkg"
	"fmt"
	"math"
	"slices"
)

func main() {

	// Time complexity: O(nlogn)
	// Space complexity: O(n)
	// Reason for time complexity:
	// 1. Read csv O(n)
	// 2. Split the list O(n)
	// 3. Sort list slices.Sort uses a variant quicksort O(nlogN)
	// 4. Loop through the list O(n)
	// Overall O(nlog(n)

	list := pkg.MustReadCSV("day1/input.csv")
	list1 := []int{}
	list2 := []int{}
	for _, numbers := range list {
		list1 = append(list1, numbers[0])
		list2 = append(list2, numbers[1])
	}

	fmt.Println(calculateDiff(list1, list2))
}

func calculateDiff(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	totalDiff := 0
	for i := range list1 {
		number1 := list1[i]
		number2 := list2[i]
		diff := int(math.Abs(float64(number2 - number1)))

		totalDiff += diff
	}
	return totalDiff
}
