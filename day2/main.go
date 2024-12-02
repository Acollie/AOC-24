package main

import (
	"AOC-24/pkg"
	"fmt"
)

func main() {
	inputString := pkg.MustReadTxt("day2/input.txt")
	var input [][]int
	for _, line := range inputString {
		var numbers []int
		for _, value := range line {
			number := 0
			fmt.Sscanf(value, "%d", &number)
			numbers = append(numbers, number)
		}
		input = append(input, numbers)
	}
	totalSafe := 0
	for _, list := range input {
		report := calculateReport(list)
		if report == Safe {
			totalSafe++
		}
	}
	fmt.Println(totalSafe)

}

type Report string

const (
	Safe   Report = "Safe"
	Unsafe Report = "Unsafe"
)

func calculateReport(list []int) Report {
	// The levels are either all increasing or all decreasing
	// Any two adjacent levels differ by at least one and at most three
	isDecreasing := false
	isIncreasing := false
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			isDecreasing = true
		}
		if list[i] < list[i+1] {
			isIncreasing = true
		}
		if isIncreasing && isDecreasing {
			return Unsafe
		}
		if list[i] == list[i+1] {
			return Unsafe
		}
		if list[i] > list[i+1]+3 {
			return Unsafe
		} else if list[i]+3 < list[i+1] {
			return Unsafe
		}
	}
	if !isDecreasing && !isIncreasing {
		return Unsafe
	}

	return Safe
}
