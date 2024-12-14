package day14

import (
	"aoc24/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type vec2d struct {
	x, y int
}

func (v *vec2d) add(otherV *vec2d) vec2d {
	return vec2d{v.x + otherV.x, v.y + otherV.y}
}

func (v *vec2d) sub(otherV *vec2d) vec2d {
	return vec2d{v.x - otherV.x, v.y - otherV.y}
}

func (v *vec2d) lenSquared() int {
	return v.x*v.x + v.y*v.y
}

type robot struct {
	pos vec2d
	vel vec2d
}

func parseRobots(fileName string) []robot {
	var robots []robot = nil

	utils.ReadFileWithParse(fileName, func(line string) {
		vals := strings.Split(line, " ")
		if len(vals) != 2 {
			panic("What?!")
		}

		posStr := strings.TrimPrefix(vals[0], "p=")
		posValStr := strings.Split(posStr, ",")
		if len(posValStr) != 2 {
			panic("What?!")
		}

		posX, err := strconv.Atoi(posValStr[0])
		if err != nil {
			panic(err)
		}

		posY, err := strconv.Atoi(posValStr[1])
		if err != nil {
			panic(err)
		}

		velStr := strings.TrimPrefix(vals[1], "v=")
		velValStr := strings.Split(velStr, ",")
		if len(velValStr) != 2 {
			panic("What?!")
		}

		velX, err := strconv.Atoi(velValStr[0])
		if err != nil {
			panic(err)
		}

		velY, err := strconv.Atoi(velValStr[1])
		if err != nil {
			panic(err)
		}

		r := robot{vec2d{posX, posY}, vec2d{velX, velY}}

		robots = append(robots, r)
	})

	return robots
}

func simulateRobot(r *robot, gridWidth, gridHeight int) {
	r.pos = r.pos.add(&r.vel)

	//Wrap around, if required.
	if r.pos.x < 0 {
		r.pos.x = gridWidth + r.pos.x
	} else if r.pos.x >= gridWidth {
		r.pos.x -= gridWidth
	}

	if r.pos.y < 0 {
		r.pos.y = gridHeight + r.pos.y
	} else if r.pos.y >= gridHeight {
		r.pos.y -= gridHeight
	}
}

func printPossibleTree(robots []robot, gridWidth, gridHeight int) {
	robotPos := make([][]int, gridHeight)

	for y := 0; y < gridHeight; y++ {
		robotPos[y] = make([]int, gridWidth)
	}

	for _, r := range robots {
		robotPos[r.pos.y][r.pos.x]++
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if robotPos[y][x] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Print("\n")
	}
}

func runPuzzle1(fileName string) {
	robots := parseRobots(fileName)

	gridWidth := 101
	gridHeight := 103

	for i := 0; i < 100; i++ {
		for r := range robots {
			simulateRobot(&robots[r], gridWidth, gridHeight)
		}
	}

	quadrantWidth := gridWidth / 2
	quadrantHeight := gridHeight / 2

	topLeftCount := 0
	topRightCount := 0
	bottomLeftCount := 0
	bottomRightCount := 0

	for _, r := range robots {
		inTopHalf := r.pos.y >= 0 && r.pos.y < quadrantHeight
		inBottomHalf := r.pos.y > quadrantHeight && r.pos.y < gridHeight
		inLeftHalf := r.pos.x >= 0 && r.pos.x < quadrantWidth
		inRightHalf := r.pos.x > quadrantWidth && r.pos.x < gridWidth

		if inTopHalf && inLeftHalf {
			//Top left.
			topLeftCount++
		} else if inTopHalf && inRightHalf {
			//Top right.
			topRightCount++
		} else if inBottomHalf && inLeftHalf {
			//Bottom left.
			bottomLeftCount++
		} else if inBottomHalf && inRightHalf {
			//Bottom right.
			bottomRightCount++
		} else {
			//On edge.
		}
	}

	safetyFactor := topLeftCount * topRightCount * bottomLeftCount * bottomRightCount

	fmt.Println("Safety Factor:", safetyFactor)
}

func runPuzzle2(fileName string) {
	robots := parseRobots(fileName)

	gridWidth := 101
	gridHeight := 103

	disorderedAvgDistSqu := 0.0
	diffTolerance := 0.0

	for i := 0; i < 10000; i++ {
		for r := range robots {
			simulateRobot(&robots[r], gridWidth, gridHeight)
		}

		//I've no idea if this is the right way to do this but given that the positions
		//of the robots is random I'm calculating the average distance between all of the
		//robots. When they close together and form an image it will bring the average down.
		totalDistSqu := 0

		for i := 0; i < len(robots); i++ {
			for j := i + 1; j < len(robots); j++ {
				diff := robots[j].pos.sub(&robots[i].pos)
				distSqu := diff.lenSquared()
				totalDistSqu += distSqu
			}
		}

		avgDistSqu := float64(totalDistSqu) / float64(len(robots)-1)

		if i == 0 {
			disorderedAvgDistSqu = avgDistSqu
			diffTolerance = disorderedAvgDistSqu * 0.4
		} else {
			diff := math.Abs(disorderedAvgDistSqu - avgDistSqu)
			if diff > diffTolerance {
				fmt.Println("Is this a tree at", i+1, "...")
				printPossibleTree(robots, gridWidth, gridHeight)
			}
		}
	}
}

func Run(inputFileName string) {
	fmt.Println("Running Day 14 Puzzle 1...")
	runPuzzle1(inputFileName)

	fmt.Println("Running Day 14 Puzzle 2...")
	runPuzzle2(inputFileName)
}
