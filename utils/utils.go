package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Read2DNumFile(fileName string, delim string) [][]int {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var retNums [][]int

	for i := 0; scanner.Scan(); i++ {
		retNums = append(retNums, nil)

		nums := strings.Split(scanner.Text(), delim)

		for _, num := range nums {
			v, err := strconv.Atoi(num)

			if err != nil {
				panic(err)
			}

			retNums[i] = append(retNums[i], v)
		}
	}

	return retNums
}
