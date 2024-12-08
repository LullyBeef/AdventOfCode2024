package day5

import (
	"aoc24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type pagePair struct {
	x, y int
}

// File parsing is very specific to this day.
func parseInput(fileName string) (pagePairs []pagePair, updates [][]int) {
	readingPagePairs := true
	updateIdx := 0
	utils.ReadFileWithParse(fileName, func(line string) {
		if readingPagePairs {
			if len(line) == 0 {
				readingPagePairs = false
			} else {
				pages := strings.Split(line, "|")
				if len(pages) != 2 {
					panic("Parsing error!")
				}

				x, err := strconv.Atoi(pages[0])
				if err != nil {
					panic(err)
				}

				y, err := strconv.Atoi(pages[1])
				if err != nil {
					panic(err)
				}

				pagePairs = append(pagePairs, pagePair{x, y})
			}
		} else {
			updates = append(updates, nil)
			pageNums := strings.Split(line, ",")

			for _, num := range pageNums {
				v, err := strconv.Atoi(num)

				if err != nil {
					panic(err)
				}

				updates[updateIdx] = append(updates[updateIdx], v)
			}

			updateIdx++
		}
	})

	return
}

func checkInOrder(update []int, pagePairs []pagePair) bool {
	for _, checkPair := range pagePairs {
		xIdx := slices.Index(update, checkPair.x)
		yIdx := slices.Index(update, checkPair.y)

		if xIdx >= 0 && yIdx >= 0 && xIdx > yIdx {
			return false
		}
	}

	return true
}

func runPuzzle2(fileName string) {
	pagePairs, updates := parseInput(fileName)

	sumMiddlePages := 0
	for _, update := range updates {

		inOrder := checkInOrder(update, pagePairs)

		if !inOrder {
			//Sort them based on the rule list.
			slices.SortFunc(update, func(x, y int) int {
				for _, p := range pagePairs {
					if (x == p.x && y == p.y) || (x == p.y && x == p.x) {
						if x == p.x {
							return -1
						} else {
							return 1
						}
					}
				}
				return 0
			})

			if !checkInOrder(update, pagePairs) {
				panic("They should be sorted now!")
			}

			updateLen := len(update)
			middlePage := 0
			if (updateLen % 2) == 1 {
				middlePage = update[updateLen/2]
				sumMiddlePages += middlePage
			} else {
				panic("Can't find middle!")
			}
		}
	}

	fmt.Println("Sum Middle Pages: ", sumMiddlePages)
}

func runPuzzle1(fileName string) {
	pagePairs, updates := parseInput(fileName)

	sumMiddlePages := 0
	for _, update := range updates {
		inOrder := checkInOrder(update, pagePairs)

		if inOrder {
			updateLen := len(update)
			middlePage := 0
			if (updateLen % 2) == 1 {
				middlePage = update[updateLen/2]
				sumMiddlePages += middlePage
			} else {
				panic("Can't find middle!")
			}
		}
	}

	fmt.Println("Sum Middle Pages: ", sumMiddlePages)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 5 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 5 Puzzle 2...")
	runPuzzle2(inputFileName)
}
