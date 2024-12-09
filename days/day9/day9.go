package day9

import (
	"aoc24/utils"
	"fmt"
)

const emptyBlock int = -1

func buildDiskMap(input []int) []int {
	var diskMap []int

	for i, v := range input {
		isFreeSpace := (i % 2) == 1 //File or free space
		id := i / 2

		for vCnt := 0; vCnt < v; vCnt++ {
			if isFreeSpace {
				diskMap = append(diskMap, emptyBlock)
			} else {
				diskMap = append(diskMap, id)
			}
		}
	}

	return diskMap
}

func packMemFragment(diskMap []int) {
	endIdx := len(diskMap) - 1

	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == emptyBlock {
			for diskMap[endIdx] == emptyBlock {
				endIdx--
			}

			if endIdx > i {
				diskMap[i] = diskMap[endIdx]
				diskMap[endIdx] = emptyBlock
			} else {
				break
			}
		}
	}
}

func tryMove(diskMap []int, fromIdx int, size int) {
	foundSpaceIdx := -1

	if fromIdx >= 0 && fromIdx < len(diskMap) && diskMap[fromIdx] != emptyBlock {
		currEmptySize := 0
		currEmptyIdx := 0
		for i := 0; i < fromIdx; i++ {
			if diskMap[i] == emptyBlock {
				if currEmptySize == 0 {
					currEmptyIdx = i
				}
				currEmptySize++
			} else if currEmptySize >= size {
				foundSpaceIdx = currEmptyIdx
				break
			} else {
				currEmptySize = 0
			}
		}

		if currEmptySize >= size {
			foundSpaceIdx = currEmptyIdx
		}
	}

	if foundSpaceIdx >= 0 {
		for i := 0; i < size; i++ {
			diskMap[foundSpaceIdx+i] = diskMap[fromIdx+i]
			diskMap[fromIdx+i] = emptyBlock
		}
	}
}

func packMem(diskMap []int) {
	currFileSize := 0
	currBlock := emptyBlock
	for endIdx := len(diskMap) - 1; endIdx >= 0; endIdx-- {
		if diskMap[endIdx] != currBlock {
			tryMove(diskMap, endIdx+1, currFileSize)

			currBlock = diskMap[endIdx]
			currFileSize = 1
		} else {
			currFileSize++
		}
	}

	tryMove(diskMap, 0, currFileSize)
}

func findChecksum(diskMap []int) int {
	checkSum := 0
	for i, v := range diskMap {
		if v != emptyBlock {
			checkSum += i * v
		}
	}

	return checkSum
}

func runPuzzle(fileName string, allowFragmentation bool) {
	input := utils.Read2DNumFile(fileName, "")

	if len(input) != 1 {
		panic("Parse error!")
	}

	diskMap := buildDiskMap(input[0])
	if allowFragmentation {
		packMemFragment(diskMap)
	} else {
		packMem(diskMap)
	}
	checkSum := findChecksum(diskMap)

	fmt.Println("Checksum: ", checkSum)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 9 Puzzle 1...")
	runPuzzle(inputFileName, true)

	fmt.Println("Running Day 9 Puzzle 2...")
	runPuzzle(inputFileName, false)
}
