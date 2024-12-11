package day11

import (
	"aoc24/utils"
	"fmt"
)

func numDigits(v int) int {
	if v == 0 {
		return 1
	}

	n := 0
	temp := v
	for temp > 0 {
		temp /= 10
		n++
	}

	return n
}

func splitNum(v int, numDig int) (int, int) {
	d := 1
	for i := 0; i < numDig/2; i++ {
		d *= 10
	}

	l := v / d
	r := v % d

	return l, r
}

func runBlinks(stones []int, numBlinks int) int {
	numStones := 0

	stonesMap := make(map[int]int)

	//Initial state from input.
	for _, v := range stones {
		stonesMap[v]++
	}

	for b := 0; b < numBlinks; b++ {
		updateMap := make(map[int]int)

		for i, n := range stonesMap {
			if n > 0 {
				if i == 0 {
					updateMap[1] += stonesMap[0]
				} else {
					numDig := numDigits(i)
					numDigEven := (numDig % 2) == 0

					if numDigEven {
						l, r := splitNum(i, numDig)
						updateMap[l] += stonesMap[i]
						updateMap[r] += stonesMap[i]
					} else {
						mulIdx := i * 2024
						updateMap[mulIdx] += stonesMap[i]
					}
				}
			}
		}

		numStones = 0
		//Apply back to stones.
		for k := range stonesMap {
			delete(stonesMap, k)
		}
		for k, v := range updateMap {
			stonesMap[k] = v
			numStones += v
		}
	}

	return numStones
}

func runPuzzle(fileName string, numTimes int) {
	input := utils.Read2DNumFile(fileName, " ")

	if len(input) != 1 {
		panic("Parse error!")
	}

	numStones := runBlinks(input[0], numTimes)

	fmt.Println("Num Stones:", numStones)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 11 Puzzle 1...")
	runPuzzle(inputFileName, 25)

	fmt.Println("Running Day 11 Puzzle 2...")
	runPuzzle(inputFileName, 75)
}
