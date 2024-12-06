package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Allow custom parsing for each line.
func ReadFileWithParse(fileName string, parseFunc func(line string)) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parseFunc(scanner.Text())
	}
}

// Versions for common parsing options.
func ReadFile(fileName string) []string {
	var lines []string

	ReadFileWithParse(fileName, func(line string) {
		lines = append(lines, line)
	})

	return lines
}

func Read2DNumFile(fileName string, delim string) [][]int {
	var retNums [][]int

	lineIdx := 0
	ReadFileWithParse(fileName, func(line string) {
		retNums = append(retNums, nil)

		nums := strings.Split(line, delim)

		for _, num := range nums {
			v, err := strconv.Atoi(num)

			if err != nil {
				panic(err)
			}

			retNums[lineIdx] = append(retNums[lineIdx], v)
		}

		lineIdx++
	})

	return retNums
}
