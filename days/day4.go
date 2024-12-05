package days

import (
	"aoc24/utils"
	"fmt"
)

type Direction int

const (
	DIR_UP Direction = iota
	DIR_DOWN
	DIR_LEFT
	DIR_RIGHT
	DIR_UP_RIGHT
	DIR_UP_LEFT
	DIR_DOWN_RIGHT
	DIR_DOWN_LEFT
	DIR_MAX
)

type indexPair struct {
	a, b int
}

func findXmas(input []string, xLoc indexPair, foundChan chan int) {
	dirsFound := 0

	startRow := xLoc.a
	startCol := xLoc.b

	numRows := len(input)
	numCols := len(input[0])

	for dir := DIR_UP; dir < DIR_MAX; dir++ {
		findXmasDirFunc := func(rowInc, colInc int) {
			if input[startRow+rowInc*1][startCol+colInc*1] == 'M' &&
				input[startRow+rowInc*2][startCol+colInc*2] == 'A' &&
				input[startRow+rowInc*3][startCol+colInc*3] == 'S' {
				dirsFound++
			}
		}

		switch dir {
		case DIR_UP:
			if startRow >= 3 {
				findXmasDirFunc(-1, 0)
			}

		case DIR_DOWN:
			if startRow < numRows-3 {
				findXmasDirFunc(1, 0)
			}

		case DIR_LEFT:
			if startCol >= 3 {
				findXmasDirFunc(0, -1)
			}

		case DIR_RIGHT:
			if startCol < numCols-3 {
				findXmasDirFunc(0, 1)
			}

		case DIR_UP_RIGHT:
			if startRow >= 3 && startCol < numCols-3 {
				findXmasDirFunc(-1, 1)
			}
		case DIR_UP_LEFT:
			if startRow >= 3 && startCol >= 3 {
				findXmasDirFunc(-1, -1)
			}

		case DIR_DOWN_RIGHT:
			if startRow < numRows-3 && startCol < numCols-3 {
				findXmasDirFunc(1, 1)
			}

		case DIR_DOWN_LEFT:
			if startRow < numRows-3 && startCol >= 3 {
				findXmasDirFunc(1, -1)
			}
		}
	}

	foundChan <- dirsFound
}

func runDay4Puzzle1(fileName string) {
	input := utils.ReadFile(fileName)

	var xLocs []indexPair

	for i, line := range input {
		for j, c := range line {
			if c == 'X' {
				xLocs = append(xLocs, indexPair{i, j})
			}
		}
	}

	numXs := len(xLocs)
	foundChan := make(chan int)
	for _, xLoc := range xLocs {
		go findXmas(input, xLoc, foundChan)
	}

	numFound := 0
	for i := 0; i < numXs; i++ {
		numFound += <-foundChan
	}

	fmt.Println("Num Found: ", numFound)
}

func findMasX(input []string, aLoc indexPair, foundChan chan bool) {
	validX := false

	centreRow := aLoc.a
	centreCol := aLoc.b

	checkDirFunc := func(dir Direction) bool {
		switch dir {
		case DIR_DOWN_RIGHT:
			return input[centreRow-1][centreCol-1] == 'M' && input[centreRow+1][centreCol+1] == 'S'
		case DIR_DOWN_LEFT:
			return input[centreRow-1][centreCol+1] == 'M' && input[centreRow+1][centreCol-1] == 'S'
		case DIR_UP_RIGHT:
			return input[centreRow+1][centreCol-1] == 'M' && input[centreRow-1][centreCol+1] == 'S'
		case DIR_UP_LEFT:
			return input[centreRow+1][centreCol+1] == 'M' && input[centreRow-1][centreCol-1] == 'S'
		}

		panic("Shouldn't get here!")
	}

	//This should be redundant as we validate the index pair before kicking off the goroutines.
	validCentre := centreRow > 0 && centreRow < len(input)-1 && centreCol > 0 && centreCol < len(input[0])-1
	if validCentre {
		if checkDirFunc(DIR_DOWN_RIGHT) {
			validX = checkDirFunc(DIR_UP_RIGHT) || checkDirFunc(DIR_DOWN_LEFT)
		} else if checkDirFunc(DIR_DOWN_LEFT) {
			validX = checkDirFunc(DIR_UP_LEFT) || checkDirFunc(DIR_DOWN_RIGHT)
		} else if checkDirFunc(DIR_UP_RIGHT) {
			validX = checkDirFunc(DIR_UP_LEFT) || checkDirFunc(DIR_DOWN_RIGHT)
		} else if checkDirFunc(DIR_UP_LEFT) {
			validX = checkDirFunc(DIR_UP_RIGHT) || checkDirFunc(DIR_DOWN_LEFT)
		}
	} else {
		panic("How did this happen?!")
	}

	foundChan <- validX
}

func runDay4Puzzle2(fileName string) {
	input := utils.ReadFile(fileName)

	var aLocs []indexPair

	//In this case the center A will always have to be at least one away from the edge
	//so we can do the filtering here.
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] == 'A' {
				aLocs = append(aLocs, indexPair{i, j})
			}
		}
	}

	numAs := len(aLocs)
	foundChan := make(chan bool)
	for _, aLoc := range aLocs {
		go findMasX(input, aLoc, foundChan)
	}

	numFound := 0
	for i := 0; i < numAs; i++ {
		validX := <-foundChan

		if validX {
			numFound++
		}
	}

	fmt.Println("Num Found: ", numFound)
}

func Day4() {
	fmt.Println("Running Day 4 Puzzle 1 Example...")
	runDay4Puzzle1("inputs/day4_example.txt")

	fmt.Println("Running Day 4 Puzzle 1...")
	runDay4Puzzle1("inputs/day4.txt")

	fmt.Println("Running Day 4 Puzzle 2 Example...")
	runDay4Puzzle2("inputs/day4_example.txt")

	fmt.Println("Running Day 4 Puzzle 2...")
	runDay4Puzzle2("inputs/day4.txt")
}
