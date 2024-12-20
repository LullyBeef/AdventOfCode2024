package day18

import (
	"aoc24/utils"
	"fmt"
)

type direction int

const (
	dirRight direction = iota
	dirDown
	dirLeft
	dirUp
	dirMax
)

var dirVecs = []vec2d{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type memoryCell bool

type memoryGrid [][]memoryCell

type vec2d struct {
	x, y int
}

func (v vec2d) add(otherV vec2d) vec2d {
	return vec2d{v.x + otherV.x, v.y + otherV.y}
}

type corruptionList []vec2d

type traverseState struct {
	pos   vec2d
	dir   direction
	score int
}

func findCheapestRoute(memory memoryGrid, source, exit vec2d) int {
	mazeHeight := len(memory)
	mazeWidth := len(memory[0])

	var queue priorityQueue

	//For each position map the score of traversing each direction.
	visited := make(map[vec2d]map[direction]int)

	for dir := dirRight; dir < dirMax; dir++ {
		queue = append(queue, traverseState{source, dir, 0})
		visited[source] = map[direction]int{dir: 0}
	}

	//Dijkstra
	for len(queue) > 0 {
		current := queue.pop()

		if current.pos == exit {
			return current.score
		}

		for dir := dirRight; dir < dirMax; dir++ {
			neighbour := current.pos.add(dirVecs[dir])

			if neighbour.x >= 0 && neighbour.x < mazeWidth &&
				neighbour.y >= 0 && neighbour.y < mazeHeight &&
				!memory[neighbour.y][neighbour.x] {

				newScore := current.score + 1

				if _, exists := visited[neighbour]; !exists {
					visited[neighbour] = make(map[direction]int)
				}

				if prevScore, exists := visited[neighbour][dir]; !exists || newScore < prevScore {
					visited[neighbour][dir] = newScore
					queue.push(traverseState{neighbour, dir, newScore})
				}
			}
		}
	}

	return -1
}

type priorityQueue []traverseState

func (queue *priorityQueue) push(state traverseState) {
	*queue = append(*queue, state)

	//Keep the queue sorted by score.
	for i := len(*queue) - 1; i > 0; i-- {
		if (*queue)[i].score < (*queue)[i-1].score {
			(*queue)[i], (*queue)[i-1] = (*queue)[i-1], (*queue)[i]
		} else {
			break
		}
	}
}

func (queue *priorityQueue) pop() traverseState {
	state := (*queue)[0]
	*queue = (*queue)[1:]
	return state
}

func parseCorruptionList(fileName string) corruptionList {

	var corruptions corruptionList

	corruptionLocs := utils.Read2DNumFile(fileName, ",")

	for _, corruptionLoc := range corruptionLocs {
		if len(corruptionLoc) != 2 {
			panic("Parse error!")
		}

		corruptions = append(corruptions, vec2d{corruptionLoc[0], corruptionLoc[1]})
	}

	return corruptions
}

func makeMemory(memoryWidth int) memoryGrid {
	memory := make(memoryGrid, memoryWidth)
	for i := range memory {
		memory[i] = make([]memoryCell, memoryWidth)
	}

	return memory
}

/*func drawMemory(memory memoryGrid) {
	for y := 0; y < len(memory); y++ {
		for x := 0; x < len(memory[y]); x++ {
			if memory[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}*/

func runPuzzle1(fileName string) {
	corruptions := parseCorruptionList(fileName)

	//Example
	//memoryWidth := 7
	//simulateBytes := 12

	//Full
	memoryWidth := 71
	simulateBytes := 1024

	memory := makeMemory(memoryWidth)

	for i := 0; i < simulateBytes; i++ {
		corruption := corruptions[i]
		memory[corruption.y][corruption.x] = true
	}

	source := vec2d{0, 0}
	exit := vec2d{memoryWidth - 1, memoryWidth - 1}

	score := findCheapestRoute(memory, source, exit)

	fmt.Println("Score:", score)
}

func runPuzzle2(fileName string) {
	corruptions := parseCorruptionList(fileName)

	//Example
	//memoryWidth := 7

	//Full
	memoryWidth := 71

	memory := makeMemory(memoryWidth)

	for i := 0; i < len(corruptions); i++ {
		corruption := corruptions[i]
		memory[corruption.y][corruption.x] = true

		source := vec2d{0, 0}
		exit := vec2d{memoryWidth - 1, memoryWidth - 1}

		score := findCheapestRoute(memory, source, exit)

		if score == -1 {
			fmt.Printf("Blocked route after %d,%d!\n", corruption.x, corruption.y)
			break
		}
	}

	fmt.Println("Couldn't find a blocked route!")
}

func Run(inputFileName string) {
	fmt.Println("Running Day 18 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 18 Puzzle 2...")
	runPuzzle2(inputFileName)
}
