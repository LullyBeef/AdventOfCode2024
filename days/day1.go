package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fetchNumLists(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftNums []int
	var rightNums []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")

		if len(nums) != 2 {
			panic("Expected 2 numbers! Parsing error...")
		}

		v, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		leftNums = append(leftNums, v)

		v, err = strconv.Atoi(nums[1])

		rightNums = append(rightNums, v)
	}

	return leftNums, rightNums
}

func runDay1Puzzle1(fileName string) {
	leftNums, rightNums := fetchNumLists(fileName)

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	sumDiffs := 0

	//leftNums and rightNums will be the same size.
	for i := range leftNums {
		leftNum := leftNums[i]
		rightNum := rightNums[i]

		diff := 0
		if rightNum > leftNum {
			diff = rightNum - leftNum
		} else {
			diff = leftNum - rightNum
		}

		sumDiffs += diff
	}

	fmt.Println("sumDiffs: ", sumDiffs)
}

func runDay1Puzzle2(fileName string) {
	leftNums, rightNums := fetchNumLists(fileName)

	score := 0
	for _, leftNum := range leftNums {
		count := 0
		for _, rightNum := range rightNums {
			if leftNum == rightNum {
				count++
			}
		}

		score += leftNum * count
	}

	fmt.Println("score: ", score)
}

func Day1() {
	fmt.Println("Running Day 1 Puzzle 1 Example...")
	runDay1Puzzle1("inputs/day1_example.txt")

	fmt.Println("Running Day 1 Puzzle 1...")
	runDay1Puzzle1("inputs/day1.txt")

	fmt.Println("Running Day 1 Puzzle 2 Example...")
	runDay1Puzzle2("inputs/day1_example.txt")

	fmt.Println("Running Day 1 Puzzle 2...")
	runDay1Puzzle2("inputs/day1.txt")
}
