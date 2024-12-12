package day12

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

type region struct {
	plantType byte
	plantLocs []loc
}

func traverseRegion(grid []string, handled [][]bool, x, y int, newRegion *region) {
	if !handled[y][x] && grid[y][x] == newRegion.plantType {
		gridLen := len(grid)
		handled[y][x] = true
		newRegion.plantLocs = append(newRegion.plantLocs, loc{x, y})
		for dir := dirUp; dir < dirMax; dir++ {
			switch dir {
			case dirUp:
				if y > 0 {
					traverseRegion(grid, handled, x, y-1, newRegion)
				}
			case dirRight:
				if x < gridLen-1 {
					traverseRegion(grid, handled, x+1, y, newRegion)
				}
			case dirDown:
				if y < gridLen-1 {
					traverseRegion(grid, handled, x, y+1, newRegion)
				}
			case dirLeft:
				if x > 0 {
					traverseRegion(grid, handled, x-1, y, newRegion)
				}
			}
		}
	}
}

func findRegions(grid []string) *[]region {
	var handled [][]bool
	for y := range grid {
		handled = append(handled, nil)
		for x := 0; x < len(grid[y]); x++ {
			handled[y] = append(handled[y], false)
		}
	}

	var regions []region

	for y := range grid {
		for x := range grid[y] {
			if !handled[y][x] {
				var newRegion region
				newRegion.plantType = grid[y][x]
				traverseRegion(grid, handled, x, y, &newRegion)

				regions = append(regions, newRegion)
			}
		}
	}

	return &regions
}

func calcPerimeter(r *region) int {
	perim := len(r.plantLocs) * 4
	for _, p := range r.plantLocs {
		for dir := dirUp; dir < dirMax; dir++ {
			switch dir {
			case dirUp:
				if slices.Contains(r.plantLocs, loc{p.x, p.y - 1}) {
					perim--
				}
			case dirRight:
				if slices.Contains(r.plantLocs, loc{p.x + 1, p.y}) {
					perim--
				}
			case dirDown:
				if slices.Contains(r.plantLocs, loc{p.x, p.y + 1}) {
					perim--
				}
			case dirLeft:
				if slices.Contains(r.plantLocs, loc{p.x - 1, p.y}) {
					perim--
				}
			}
		}
	}

	return perim
}

func calcEdges(r *region) int {
	//Not really happy with this... I should probably look into edge detection algorithms...
	edgePointMap := make(map[loc][]loc)

	edgePointOffsets := [4][2]int{
		{0, 0}, //Top left
		{1, 0}, //Top right
		{0, 1}, //Bottom left
		{1, 1}, //Bottom right
	}

	for _, p := range r.plantLocs {
		for i := 0; i < 4; i++ {
			edgePoint := loc{p.x + edgePointOffsets[i][0], p.y + edgePointOffsets[i][1]}
			edgePointMap[edgePoint] = append(edgePointMap[edgePoint], p)
		}
	}

	//Count edges by ignoring any shared points.
	numEdges := 0
	for _, e := range edgePointMap {
		numPoints := len(e)
		if numPoints == 1 || numPoints == 3 { //If they're on a corner, they're an edge point.
			numEdges++
		} else if numPoints == 2 {
			//Special case for internal diagonals... This feels like bullshit...
			if e[0].x != e[1].x && e[0].y != e[1].y {
				numEdges += 2
			}
		}
	}

	return numEdges
}

func runPuzzle(fileName string, perimOrEdges bool) {
	grid := utils.ReadFile(fileName)

	regions := findRegions(grid)

	sumPrice := 0

	for i := range *regions {
		area := len((*regions)[i].plantLocs)
		price := 0

		if perimOrEdges {
			perim := calcPerimeter(&(*regions)[i])
			price = area * perim
		} else {
			edges := calcEdges(&(*regions)[i])
			price = area * edges
		}

		sumPrice += price
	}

	fmt.Println("Sum Price:", sumPrice)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 12 Puzzle 1...")
	runPuzzle(inputFileName, true)

	fmt.Println("Running Day 12 Puzzle 2...")
	runPuzzle(inputFileName, false)
}
