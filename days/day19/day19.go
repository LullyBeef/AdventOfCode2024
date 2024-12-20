package day19

import (
	"aoc24/utils"
	"fmt"
	"strconv"
	"strings"
)

type colour int

const (
	colWhite colour = iota
	colBlue
	colBlack
	colRed
	colGreen
)

type towel []colour
type towelList []towel

type pattern []colour

func (p pattern) StartsWith(t towel) bool {
	if len(p) >= len(t) {
		for i := range t {
			if p[i] != t[i] {
				return false
			}
		}
		return true
	}

	return false
}

type patternList []pattern

func (p pattern) hash() string {
	/*hash := 0

	for i, v := range p {
		hash |= (int(v) << (i * 3))
	}

	return hash*/

	//I should've just kept these as strings when parsing because this is crap.
	//I need a unique hash key for memoization, though. I was previously using
	//the above but that won't work for the full input as it exceeds int size.
	var hashStr string

	for _, v := range p {
		hashStr += strconv.Itoa(int(v))
	}

	return hashStr
}

var substrings = []string{}
var targets = []string{}

func parseInput(fileName string) (towelList, patternList) {
	readingTowels := true

	var towels towelList
	var patterns patternList

	utils.ReadFileWithParse(fileName, func(line string) {
		if len(line) > 0 {
			if readingTowels {
				towelStrs := strings.Split(line, ", ")

				for _, towelStr := range towelStrs {
					substrings = append(substrings, towelStr)

					var t towel
					for _, c := range towelStr {
						var col colour
						switch c {
						case 'w':
							col = colWhite
						case 'u':
							col = colBlue
						case 'b':
							col = colBlack
						case 'r':
							col = colRed
						case 'g':
							col = colGreen
						default:
							panic("Unhandled!")
						}

						t = append(t, col)
					}

					towels = append(towels, t)
				}

				readingTowels = false
			} else {
				targets = append(targets, line)
				var p pattern
				for _, c := range line {
					var col colour
					switch c {
					case 'w':
						col = colWhite
					case 'u':
						col = colBlue
					case 'b':
						col = colBlack
					case 'r':
						col = colRed
					case 'g':
						col = colGreen
					default:
						panic("Unhandled!")
					}

					p = append(p, col)
				}

				patterns = append(patterns, p)
			}
		}
	})

	return towels, patterns
}

func canMakePattern(p pattern, towels towelList, cache map[string]int) int {
	if len(p) == 0 {
		return 1
	}

	if v, exists := cache[p.hash()]; exists {
		return v
	}

	numWays := 0
	for _, t := range towels {
		if p.StartsWith(t) {
			numWays += canMakePattern(p[len(t):], towels, cache)
		}
	}

	cache[p.hash()] = numWays
	return numWays
}

func runPuzzle1(fileName string) {
	towels, patterns := parseInput(fileName)

	numPossible := 0

	for _, p := range patterns {
		cache := make(map[string]int)

		if canMakePattern(p, towels, cache) > 0 {
			numPossible++
		}
	}

	fmt.Println("Num Possible:", numPossible)
}

func runPuzzle2(fileName string) {
	towels, patterns := parseInput(fileName)

	totalWays := 0

	for _, p := range patterns {
		cache := make(map[string]int)

		numWays := canMakePattern(p, towels, cache)
		totalWays += numWays
	}

	fmt.Println("Total Ways:", totalWays)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 19 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 19 Puzzle 2...")
	runPuzzle2(inputFileName)
}
