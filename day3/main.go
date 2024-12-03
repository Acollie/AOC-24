package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// read "input.txt" file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file contents into a string
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the content to a string
	input := string(content)
	muls := fetchMuls(input)
	total := mutiplyMuls(muls)
	log.Printf("Total: %d", total)
}

func mutiplyMuls(input [][]int) int {
	total := 0
	for _, mul := range input {
		row := mul[0] * mul[1]
		total += row
	}
	return total
}

func fetchMuls(input string) [][]int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	var result [][]int
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		row := []int{first, second}
		result = append(result, row)
	}
	return result
}
