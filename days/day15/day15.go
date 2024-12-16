package day15

import (
	"aoc24/utils"
	"fmt"
)

type direction int

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft
	dirMax
)

var dirVecs = []vec2d{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type blockType int

const (
	blockEmpty blockType = iota
	blockBoxSingle
	blockBoxLeft
	blockBoxRight
	blockWall
)

type vec2d struct {
	x, y int
}

func (v *vec2d) add(otherV *vec2d) vec2d {
	return vec2d{v.x + otherV.x, v.y + otherV.y}
}

type warehouse [][]blockType
type instructionList []direction

func parseInput(fileName string, bigWarehouse bool) (*warehouse, *instructionList, *vec2d) {
	var wh warehouse
	var instructions instructionList
	var robotPos vec2d

	input := utils.ReadFile(fileName)

	parsingInstructions := false

	for y := 0; y < len(input); y++ {
		lineLen := len(input[y])

		if lineLen == 0 {
			parsingInstructions = true
			continue
		}

		if parsingInstructions {
			for _, c := range input[y] {
				var dir direction

				switch c {
				case '<':
					dir = dirLeft
				case 'v':
					dir = dirDown
				case '^':
					dir = dirUp
				case '>':
					dir = dirRight
				default:
					panic("Unhandled!")
				}

				instructions = append(instructions, dir)
			}
		} else {
			whWidth := lineLen
			if bigWarehouse {
				whWidth *= 2
			}
			whLine := make([]blockType, whWidth)

			for x := 0; x < lineLen; x++ {
				whx := x
				if bigWarehouse {
					whx *= 2
				}
				switch input[y][x] {
				case '#':
					whLine[whx] = blockWall
					if bigWarehouse {
						whLine[whx+1] = blockWall
					}
				case '.':
					whLine[whx] = blockEmpty
					if bigWarehouse {
						whLine[whx+1] = blockEmpty
					}
				case 'O':
					if bigWarehouse {
						whLine[whx] = blockBoxLeft
						whLine[whx+1] = blockBoxRight
					} else {
						whLine[whx] = blockBoxSingle
					}
				case '@':
					whLine[whx] = blockEmpty
					if bigWarehouse {
						whLine[whx+1] = blockEmpty
					}
					robotPos = vec2d{whx, y}
				default:
					panic("Unhandled!")
				}
			}

			wh = append(wh, whLine)
		}
	}

	return &wh, &instructions, &robotPos
}

/*func drawWarehouse(wh *warehouse, robotPos *vec2d) {
	for y := 0; y < len(*wh); y++ {
		for x := 0; x < len((*wh)[y]); x++ {
			if robotPos.x == x && robotPos.y == y {
				fmt.Print("@")
			} else {
				switch (*wh)[y][x] {
				case blockEmpty:
					fmt.Print(".")
				case blockBoxSingle:
					fmt.Print("O")
				case blockBoxLeft:
					fmt.Print("[")
				case blockBoxRight:
					fmt.Print("]")
				case blockWall:
					fmt.Print("#")
				default:
					panic("Unhandled!")
				}
			}
		}

		fmt.Print("\n")
	}
}*/

func canPushBlock(wh *warehouse, blockPos *vec2d, dir direction, checkSibling bool) bool {
	thisBlock := (*wh)[blockPos.y][blockPos.x]

	if thisBlock == blockWall {
		return false
	} else if thisBlock == blockBoxSingle || thisBlock == blockBoxLeft || thisBlock == blockBoxRight {
		if (dir == dirUp || dir == dirDown) && checkSibling {
			if thisBlock == blockBoxLeft {
				siblingBlockPos := blockPos.add(&dirVecs[dirRight])
				if !canPushBlock(wh, &siblingBlockPos, dir, false) {
					return false
				}
			} else if thisBlock == blockBoxRight {
				siblingBlockPos := blockPos.add(&dirVecs[dirLeft])
				if !canPushBlock(wh, &siblingBlockPos, dir, false) {
					return false
				}
			}
		}

		dirVec := &dirVecs[dir]
		nextBlockPos := blockPos.add(dirVec)
		if !canPushBlock(wh, &nextBlockPos, dir, true) {
			return false
		}
	}

	return true
}

func pushBlock(wh *warehouse, blockPos *vec2d, dir direction, pushSibling bool) bool {
	if canPushBlock(wh, blockPos, dir, pushSibling) {
		thisBlock := (*wh)[blockPos.y][blockPos.x]

		if thisBlock == blockBoxSingle || thisBlock == blockBoxLeft || thisBlock == blockBoxRight {
			if (dir == dirUp || dir == dirDown) && pushSibling {
				if thisBlock == blockBoxLeft {
					siblingBlockPos := blockPos.add(&dirVecs[dirRight])
					if !pushBlock(wh, &siblingBlockPos, dir, false) {
						return false
					}
				} else if thisBlock == blockBoxRight {
					siblingBlockPos := blockPos.add(&dirVecs[dirLeft])
					if !pushBlock(wh, &siblingBlockPos, dir, false) {
						return false
					}
				}
			}

			dirVec := &dirVecs[dir]
			nextBlockPos := blockPos.add(dirVec)
			pushBlock(wh, &nextBlockPos, dir, true)

			moveBlock := (*wh)[blockPos.y][blockPos.x]
			(*wh)[blockPos.y][blockPos.x] = blockEmpty
			(*wh)[nextBlockPos.y][nextBlockPos.x] = moveBlock

			return true
		}
	}

	return false
}

func processInstructions(wh *warehouse, instructions *instructionList, robotPos *vec2d) {
	for _, d := range *instructions {
		dirVec := &dirVecs[d]
		nextPos := robotPos.add(dirVec)
		nextBlock := (*wh)[nextPos.y][nextPos.x]

		if nextBlock == blockEmpty {
			*robotPos = nextPos
		} else if nextBlock == blockBoxSingle || nextBlock == blockBoxLeft || nextBlock == blockBoxRight {
			if pushBlock(wh, &nextPos, d, d == dirUp || d == dirDown) {
				*robotPos = nextPos
			}
		} else if nextBlock != blockWall {
			panic("What?!")
		}
	}
}

func runPuzzle1(fileName string) {
	wh, instructionList, robotPos := parseInput(fileName, false)
	processInstructions(wh, instructionList, robotPos)

	sumCoords := 0

	for y := 0; y < len(*wh); y++ {
		for x := 0; x < len((*wh)[y]); x++ {
			if (*wh)[y][x] == blockBoxSingle {
				coord := 100*y + x
				sumCoords += coord
			}
		}
	}

	fmt.Println("Sum Coords:", sumCoords)
}

func runPuzzle2(fileName string) {
	wh, instructionList, robotPos := parseInput(fileName, true)
	processInstructions(wh, instructionList, robotPos)

	sumCoords := 0

	for y := 0; y < len(*wh); y++ {
		for x := 0; x < len((*wh)[y]); x++ {
			if (*wh)[y][x] == blockBoxLeft {
				coord := 100*y + x
				sumCoords += coord
			}
		}
	}

	fmt.Println("Sum Coords:", sumCoords)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 15 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 15 Puzzle 2...")
	runPuzzle2(inputFileName)
}
