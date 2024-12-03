package days

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

func runDay3Puzzle1(fileName string) {
	input := utils.ReadFile(fileName)
	sumProducts := handleInstructionList(input, true)
	fmt.Println("Sum Products: ", sumProducts)
}

func runDay3Puzzle2(fileName string) {
	input := utils.ReadFile(fileName)
	sumProducts := handleInstructionList(input, false)
	fmt.Println("Sum Products: ", sumProducts)
}

func Day3() {
	fmt.Println("Running Day 3 Puzzle 1 Example...")
	runDay3Puzzle1("inputs/day3_example1.txt")

	fmt.Println("Running Day 3 Puzzle 1...")
	runDay3Puzzle1("inputs/day3.txt")

	fmt.Println("Running Day 3 Puzzle 2 Example...")
	runDay3Puzzle2("inputs/day3_example2.txt")

	fmt.Println("Running Day 3 Puzzle 2...")
	runDay3Puzzle2("inputs/day3.txt")
}
