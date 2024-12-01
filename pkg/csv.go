package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
)

func MustReadCSV(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var result [][]int
	for _, record := range records {
		var numbers []int
		for _, value := range record {
			number := 0
			fmt.Sscanf(value, "%d", &number)
			numbers = append(numbers, number)
		}
		result = append(result, numbers)
	}

	return result
}
