package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func MustReadTxt(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result [][]string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				result = append(result, strings.Fields(line))
			}
			break
		}
		if err != nil {
			panic(err)
		}
		result = append(result, strings.Fields(line))
	}
	return result
}
