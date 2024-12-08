package day3

import (
	"aoc24/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func handleInstructionList(input []string, onlyMul bool) int {
	mulEnabled := true
	sumProducts := 0

	for _, line := range input {
		var re *regexp.Regexp = nil

		if onlyMul {
			re = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		} else {
			re = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
		}

		instrs := re.FindAllString(line, -1)

		for _, instr := range instrs {
			switch {
			case strings.HasPrefix(instr, "mul"):
				if mulEnabled {
					mul := instr[4 : len(instr)-1]
					nums := strings.Split(mul, ",")

					if len(nums) != 2 {
						panic("Should be two numbers!")
					}

					x, err := strconv.Atoi(nums[0])

					if err != nil {
						panic(err)
					}

					y, err := strconv.Atoi(nums[1])

					if err != nil {
						panic(err)
					}

					sumProducts += x * y
				}
			case instr == "don't()":
				mulEnabled = false
			case instr == "do()":
				mulEnabled = true
			default:
				panic("Unhandled instruction!")
			}
		}
	}

	return sumProducts
}

func runPuzzle1(fileName string) {
	input := utils.ReadFile(fileName)
	sumProducts := handleInstructionList(input, true)
	fmt.Println("Sum Products: ", sumProducts)
}

func runPuzzle2(fileName string) {
	input := utils.ReadFile(fileName)
	sumProducts := handleInstructionList(input, false)
	fmt.Println("Sum Products: ", sumProducts)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 3 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 3 Puzzle 2...")
	runPuzzle2(inputFileName)
}
