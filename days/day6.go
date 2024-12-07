package days

import (
	"aoc24/utils"
	"fmt"
)

type vec2d struct {
	x, y int
}

func (v *vec2d) add(otherV *vec2d) {
	v.x += otherV.x
	v.y += otherV.y
}

func (v *vec2d) equals(otherV *vec2d) bool {
	return v.x == otherV.x && v.y == otherV.y
}

type guardDir int

const (
	guardDirUp guardDir = iota
	guardDirRight
	guardDirDown
	guardDirLeft
	guardDirMax
)

var guardDirs = []vec2d{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

// outVisited must be same size as grid.
func runPath(guardPosition vec2d, guardDirection guardDir, grid [][]bool, outVisited [][]int) int {
	lastGuardPosition := guardPosition

	outVisited[guardPosition.y][guardPosition.x] |= (1 << guardDirection)

	for {
		lastGuardPosition = guardPosition
		guardPosition.add(&guardDirs[guardDirection])

		if guardPosition.x < 0 || guardPosition.x > len(grid[0])-1 ||
			guardPosition.y < 0 || guardPosition.y > len(grid)-1 {
			break
		}

		blocked := grid[guardPosition.y][guardPosition.x]

		if blocked {
			guardPosition = lastGuardPosition
			guardDirection = (guardDirection + 1) % guardDirMax
		} else {
			if outVisited[guardPosition.y][guardPosition.x]&(1<<guardDirection) != 0 {
				return -1 //Loop
			}

			outVisited[guardPosition.y][guardPosition.x] |= (1 << guardDirection)
		}
	}

	numVisited := 0

	for y := 0; y < len(outVisited); y++ {
		for x := 0; x < len(outVisited[y]); x++ {
			if outVisited[y][x] != 0 {
				numVisited++
			}
		}
	}

	return numVisited
}

func parseInitialState(input []string) (guardPosition vec2d, guardDirection guardDir, grid [][]bool) {
	for y, line := range input {
		grid = append(grid, nil)
		for x, c := range line {
			var obj bool
			switch c {
			case '.':
				obj = false
			case '#':
				obj = true
			case '^':
				obj = false
				guardDirection = guardDirUp
				guardPosition = vec2d{x, y}
			default:
				panic("Unsupported!")
			}

			grid[y] = append(grid[y], obj)
		}
	}

	return
}

func runDay6Puzzle1(fileName string) {
	input := utils.ReadFile(fileName)
	guardPosition, guardDirection, grid := parseInitialState(input)

	var visited [][]int
	for _, g := range grid {
		visited = append(visited, make([]int, len(g)))
	}

	numVisited := runPath(guardPosition, guardDirection, grid, visited)

	if numVisited < 0 {
		panic("Invalid!")
	}

	fmt.Println("Num Visited: ", numVisited)
}

func runDay6Puzzle2(fileName string) {
	input := utils.ReadFile(fileName)
	guardPosition, guardDirection, grid := parseInitialState(input)

	var visited [][]int
	for _, g := range grid {
		visited = append(visited, make([]int, len(g)))
	}

	var newBlockPos []vec2d

	//First find initial path where we should be placing blocks.
	runPath(guardPosition, guardDirection, grid, visited)
	yLen := len(visited)
	for y := 0; y < yLen; y++ {
		xLen := len(visited[y])
		for x := 0; x < xLen; x++ {
			if visited[y][x] != 0 {
				newBlockPos = append(newBlockPos, vec2d{x, y})
			}
		}
	}

	numLoops := 0

	//Now run the check for the all the new blocks we're expecting and see how many loop...
	//We could probably slap this in a goroutine but then we'd have to copy the grid array for
	//each iteration...
	for _, b := range newBlockPos {
		//Reset visited
		for y := range visited {
			for x := range visited[y] {
				visited[y][x] = 0
			}
		}

		//Add new block
		grid[b.y][b.x] = true

		//Run check
		loops := runPath(guardPosition, guardDirection, grid, visited) == -1

		if loops {
			numLoops++
		}

		//Remove new block
		grid[b.y][b.x] = false
	}

	fmt.Println("Num Loops: ", numLoops)
}

func Day6() {
	fmt.Println("Running Day 6 Puzzle 1 Example...")
	runDay6Puzzle1("inputs/day6_example.txt")

	fmt.Println("Running Day 6 Puzzle 1...")
	runDay6Puzzle1("inputs/day6.txt")

	fmt.Println("Running Day 6 Puzzle 2 Example...")
	runDay6Puzzle2("inputs/day6_example.txt")

	fmt.Println("Running Day 6 Puzzle 2...")
	runDay6Puzzle2("inputs/day6.txt")
}
