package day16

import (
	"aoc24/utils"
	"fmt"
	"math"
	"slices"
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

type mazeCell int

type mazeGrid [][]mazeCell

const (
	cellEmpty mazeCell = iota
	cellWall
)

type vec2d struct {
	x, y int
}

func (v vec2d) add(otherV vec2d) vec2d {
	return vec2d{v.x + otherV.x, v.y + otherV.y}
}

func parseMaze(fileName string) (maze mazeGrid, source, exit vec2d) {
	input := utils.ReadFile(fileName)

	numRows := len(input)

	for y := 0; y < numRows; y++ {
		rowLen := len(input[y])
		row := make([]mazeCell, rowLen)

		for x := 0; x < rowLen; x++ {
			switch input[y][x] {
			case '#':
				row[x] = cellWall
			case '.':
				row[x] = cellEmpty
			case 'S':
				source = vec2d{x, y}
				row[x] = cellEmpty
			case 'E':
				exit = vec2d{x, y}
				row[x] = cellEmpty
			default:
				panic("Unhandled!")
			}
		}

		maze = append(maze, row)
	}

	return
}

/*func drawMaze(maze mazeGrid, source, exit vec2d, paths [][]vec2d) {
	numRows := len(maze)
	for y := 0; y < numRows; y++ {
		rowLen := len(maze[y])
		for x := 0; x < rowLen; x++ {
			cellOveridden := false
			if paths == nil {
				if source.x == x && source.y == y {
					fmt.Print("S")
					cellOveridden = true
				} else if exit.x == x && exit.y == y {
					fmt.Print("E")
					cellOveridden = true
				}
			} else {
				for _, path := range paths {
					if slices.Contains(path, vec2d{x, y}) {
						fmt.Print("O")
						cellOveridden = true
						break
					}
				}
			}

			if !cellOveridden {
				switch maze[y][x] {
				case cellEmpty:
					fmt.Print(".")
				case cellWall:
					fmt.Print("#")
				default:
					panic("Unhandled!")
				}
			}
		}

		fmt.Print("\n")
	}
}*/

type traverseState struct {
	pos   vec2d
	dir   direction
	score int
	path  []vec2d
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

func findCheapestRoutes(maze mazeGrid, source, exit vec2d) (int, [][]vec2d) {
	mazeHeight := len(maze)
	mazeWidth := len(maze[0])

	var queue priorityQueue

	//For each position map the score of traversing each direction.
	visited := make(map[vec2d]map[direction]int)

	for dir := dirRight; dir < dirMax; dir++ {
		queue = append(queue, traverseState{source, dir, 0, []vec2d{source}})
		visited[source] = map[direction]int{dir: 0}
	}

	//Store the paths for part 2
	minScore := math.MaxInt
	var paths [][]vec2d

	//Dijkstra
	for len(queue) > 0 {
		current := queue.pop()

		if current.pos == exit {
			//return current.score
			if current.score < minScore {
				minScore = current.score
				paths = [][]vec2d{current.path}
			} else if current.score == minScore {
				paths = append(paths, current.path)
			}

			continue
		}

		for dir := dirRight; dir < dirMax; dir++ {
			neighbour := current.pos.add(dirVecs[dir])

			if neighbour.x >= 0 && neighbour.x < mazeWidth &&
				neighbour.y >= 0 && neighbour.y < mazeHeight &&
				maze[neighbour.y][neighbour.x] == cellEmpty {

				newScore := current.score + 1

				//It's expensive to turn!
				if current.dir != dir {
					newScore += 1000
				}

				if _, exists := visited[neighbour]; !exists {
					visited[neighbour] = make(map[direction]int)
				}

				//Used <= to allow for paths to diverge.
				if prevScore, exists := visited[neighbour][dir]; !exists || newScore <= prevScore {
					newPath := make([]vec2d, len(current.path))
					copy(newPath, current.path)
					newPath = append(newPath, neighbour)
					visited[neighbour][dir] = newScore
					queue.push(traverseState{neighbour, dir, newScore, newPath})
				}
			}
		}
	}

	return minScore, paths
}

func runPuzzle(fileName string, scoreOrPaths bool) {
	maze, source, exit := parseMaze(fileName)
	score, paths := findCheapestRoutes(maze, source, exit)

	//drawMaze(maze, source, exit, paths)

	if scoreOrPaths {
		fmt.Println("Score:", score)
	} else {
		var usedLocs []vec2d
		for _, path := range paths {
			for _, loc := range path {
				if !slices.Contains(usedLocs, loc) {
					usedLocs = append(usedLocs, loc)
				}
			}
		}

		fmt.Println("Path:", len(usedLocs))
	}
}

func Run(inputFileName string) {
	fmt.Println("Running Day 16 Puzzle 1...")
	runPuzzle(inputFileName, true)

	fmt.Println("Running Day 16 Puzzle 2...")
	runPuzzle(inputFileName, false)
}
