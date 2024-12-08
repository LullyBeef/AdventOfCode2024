package day7

import (
	"aoc24/utils"
	"fmt"
	"strconv"
	"strings"
)

type equationOp int

const (
	opAdd equationOp = iota
	opMul
	opConcat
	opMax
)

type equation struct {
	expectedValue int
	numbers       []int
}

func parseInput(fileName string) []equation {
	var equations []equation

	utils.ReadFileWithParse(fileName, func(line string) {
		vals := strings.Split(line, ": ")

		if len(vals) != 2 {
			panic("Parse error!")
		}

		var eq equation

		numStrs := strings.Split(vals[1], " ")

		expectedVal, err := strconv.Atoi(vals[0])
		eq.expectedValue = expectedVal

		if err != nil {
			panic("Parse error!")
		}

		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)

			if err != nil {
				panic("Parse error!")
			}

			eq.numbers = append(eq.numbers, num)
		}

		equations = append(equations, eq)
	})

	return equations
}

func runPuzzle2(fileName string) {
	//I'm just going to be stubborn and store two bits per operation.
	equations := parseInput(fileName)

	sumValid := 0

	for _, eq := range equations {
		numOps := len(eq.numbers) - 1
		maxOpCount := (1 << (numOps * 2))

		for i := 0; i < maxOpCount; i++ {
			val := eq.numbers[0]
			invalidOp := false
			for j := 0; j < numOps; j++ {
				opTypeNum := (i >> (j * 2)) & 3

				if opTypeNum < int(opMax) {

					opType := equationOp(opTypeNum)
					opNum := eq.numbers[j+1]

					switch opType {
					case opAdd:
						val += opNum
					case opMul:
						val *= opNum
					case opConcat:
						mult := 1
						for i := opNum; i > 0; i /= 10 {
							mult *= 10
						}

						val = val*mult + opNum
					}

				} else {
					invalidOp = true
				}
			}

			if !invalidOp {
				if val == eq.expectedValue {
					sumValid += eq.expectedValue
					break
				}
			}
		}
	}

	fmt.Println("Sum Valid: ", sumValid)
}

func runPuzzle1(fileName string) {
	equations := parseInput(fileName)

	sumValid := 0

	//As there's only two operations we can just count up in binary and do the op
	//for each number index... Obviously this will be a problem when they inevitably
	//add more operations in puzzle 2!
	for _, eq := range equations {
		numOps := len(eq.numbers) - 1
		maxOpCount := (1 << numOps)

		for i := 0; i < maxOpCount; i++ {
			val := eq.numbers[0]

			for j := 0; j < numOps; j++ {
				opTypeNum := (i >> j) & 1
				opType := equationOp(opTypeNum)
				opNum := eq.numbers[j+1]

				switch opType {
				case opAdd:
					val += opNum

				case opMul:
					val *= opNum
				}
			}

			if val == eq.expectedValue {
				sumValid += eq.expectedValue
				break
			}
		}
	}

	fmt.Println("Sum Valid: ", sumValid)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 7 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 7 Puzzle 2...")
	runPuzzle2(inputFileName)
}
