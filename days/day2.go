package days

import (
	"aoc24/utils"
	"fmt"
)

func checkRowSafe(row []int, safeChan chan bool) {
	rowLen := len(row)
	if rowLen <= 1 {
		panic("Row too small!")
	}

	increasing := row[1] > row[0]
	safe := true
	for i := 1; i < rowLen; i++ {
		diff := 0
		inc := row[i] > row[i-1]

		if inc != increasing {
			safe = false
			break
		}

		if inc {
			diff = row[i] - row[i-1]
		} else {
			diff = row[i-1] - row[i]
		}

		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}

	safeChan <- safe
}

// Wrapper for the row slice so we don't actually have to remove an entry, for performance.
// This may have worked better as a Reader but I haven't familiarised myself with them yet...
type ignorableRow struct {
	row         []int
	ignoreIndex int
}

func makeIgnorableRow(row []int, ignoreIndex int) ignorableRow {
	return ignorableRow{row, ignoreIndex}
}

func (r ignorableRow) len() int {
	rowLen := len(r.row)
	if r.ignoreIndex >= 0 {
		rowLen--
	}
	return rowLen
}

func (r ignorableRow) at(i int) int {
	if r.ignoreIndex >= 0 && i >= r.ignoreIndex {
		i++
	}
	return r.row[i]
}

func checkRowSafeWithDampener(row []int, safeChan chan bool) {
	//First check the row with -1 so no indices are ignored. If that fails then
	//run through the array ignoring each index until we find one that's safe or
	//we've checked the whole row.
	safe := false
	for ignoreIdx := -1; ignoreIdx < len(row); ignoreIdx++ {
		ignorableRow := makeIgnorableRow(row, ignoreIdx)

		if ignorableRow.len() <= 1 {
			panic("Row too small!")
		}

		increasing := ignorableRow.at(1) > ignorableRow.at(0)
		safe = true
		for i := 1; i < ignorableRow.len(); i++ {
			diff := 0
			inc := ignorableRow.at(i) > ignorableRow.at(i-1)

			if inc != increasing {
				safe = false
				break
			}

			if inc {
				diff = ignorableRow.at(i) - ignorableRow.at(i-1)
			} else {
				diff = ignorableRow.at(i-1) - ignorableRow.at(i)
			}

			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}

		if safe {
			break
		}
	}

	safeChan <- safe
}

func runDay2Puzzle2(fileName string) {
	input := utils.Read2DNumFile(fileName, " ")

	numRows := len(input)
	safeChan := make(chan bool)
	for _, row := range input {
		go checkRowSafeWithDampener(row, safeChan)
	}

	numSafe := 0
	for i := 0; i < numRows; i++ {
		safe := <-safeChan
		if safe {
			numSafe++
		}
	}

	fmt.Println("Num Safe: ", numSafe)
}

func runDay2Puzzle1(fileName string) {
	input := utils.Read2DNumFile(fileName, " ")

	numRows := len(input)
	safeChan := make(chan bool)
	for _, row := range input {
		go checkRowSafe(row, safeChan)
	}

	numSafe := 0
	for i := 0; i < numRows; i++ {
		safe := <-safeChan
		if safe {
			numSafe++
		}
	}

	fmt.Println("Num Safe: ", numSafe)
}

func Day2() {
	fmt.Println("Running Day 2 Puzzle 1 Example...")
	runDay2Puzzle1("inputs/day2_example.txt")

	fmt.Println("Running Day 2 Puzzle 1...")
	runDay2Puzzle1("inputs/day2.txt")

	fmt.Println("Running Day 2 Puzzle 2 Example...")
	runDay2Puzzle2("inputs/day2_example.txt")

	fmt.Println("Running Day 2 Puzzle 2...")
	runDay2Puzzle2("inputs/day2.txt")
}
