package day10

import (
	"aoc24/utils"
	"fmt"
	"slices"
)

type direction int

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft
	dirMax
)

type loc struct {
	x, y int
}

func runTrailHead(grid [][]int, x, y, want int, outFoundEnds *[]loc) int {
	if grid[y][x] != want {
		return 0 //Dead end.
	} else if grid[y][x] == 9 {
		loc := loc{x, y}
		if !slices.Contains(*outFoundEnds, loc) {
			*outFoundEnds = append(*outFoundEnds, loc)
		}
		return 1
	}

	rating := 0

	//Input is a square.
	gridLen := len(grid)

	for dir := dirUp; dir < dirMax; dir++ {

		switch {
		case dir == dirUp && y > 0:
			rating += runTrailHead(grid, x, y-1, want+1, outFoundEnds)
		case dir == dirRight && x < gridLen-1:
			rating += runTrailHead(grid, x+1, y, want+1, outFoundEnds)
		case dir == dirDown && y < gridLen-1:
			rating += runTrailHead(grid, x, y+1, want+1, outFoundEnds)
		case dir == dirLeft && x > 0:
			rating += runTrailHead(grid, x-1, y, want+1, outFoundEnds)
		}
	}

	return rating
}

func runPuzzle(fileName string, scoreOrRating bool) {
	input := utils.Read2DNumFile(fileName, "")

	sumScores := 0
	sumRatings := 0

	inputLen := len(input)
	for y := 0; y < inputLen; y++ {
		for x := 0; x < inputLen; x++ {
			if input[y][x] == 0 {
				var foundEnds []loc
				rating := runTrailHead(input, x, y, 0, &foundEnds)
				score := len(foundEnds)

				sumScores += score
				sumRatings += rating
			}
		}
	}

	if scoreOrRating {
		fmt.Println("Sum Scores:", sumScores)
	} else {
		fmt.Println("Sum Ratings:", sumRatings)
	}
}

func Run(inputFileName string) {
	fmt.Println("Running Day 10 Puzzle 1...")
	runPuzzle(inputFileName, true)

	fmt.Println("Running Day 10 Puzzle 2...")
	runPuzzle(inputFileName, false)
}
