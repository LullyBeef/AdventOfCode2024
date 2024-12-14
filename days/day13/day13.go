package day13

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

type machine struct {
	buttonA  vec2d
	buttonB  vec2d
	prizeLoc vec2d
}

func parseMachines(fileName string) []machine {
	var machines []machine = nil

	var currentMachine machine

	utils.ReadFileWithParse(fileName, func(line string) {
		if line != "" {
			sections := strings.Split(line, ": ")
			switch sections[0] {
			case "Button A":
				vals := strings.Split(sections[1], ", ")
				xValStr := strings.TrimPrefix(vals[0], "X+")
				yValStr := strings.TrimPrefix(vals[1], "Y+")

				xVal, err := strconv.Atoi(xValStr)
				if err != nil {
					panic(err)
				}

				yVal, err := strconv.Atoi(yValStr)
				if err != nil {
					panic(err)
				}

				currentMachine.buttonA.x = xVal
				currentMachine.buttonA.y = yVal

			case "Button B":
				vals := strings.Split(sections[1], ", ")
				xValStr := strings.TrimPrefix(vals[0], "X+")
				yValStr := strings.TrimPrefix(vals[1], "Y+")

				xVal, err := strconv.Atoi(xValStr)
				if err != nil {
					panic(err)
				}

				yVal, err := strconv.Atoi(yValStr)
				if err != nil {
					panic(err)
				}

				currentMachine.buttonB.x = xVal
				currentMachine.buttonB.y = yVal

			case "Prize":
				vals := strings.Split(sections[1], ", ")
				xValStr := strings.TrimPrefix(vals[0], "X=")
				yValStr := strings.TrimPrefix(vals[1], "Y=")

				xVal, err := strconv.Atoi(xValStr)
				if err != nil {
					panic(err)
				}

				yVal, err := strconv.Atoi(yValStr)
				if err != nil {
					panic(err)
				}

				currentMachine.prizeLoc.x = xVal
				currentMachine.prizeLoc.y = yVal

				machines = append(machines, currentMachine)

			default:
				panic("Unhandled!")
			}
		}
	})

	return machines
}

//Old brute force method.
/*func testMachine(m *machine, maxTimes int) (bool, int, int) {
	for aCount := 0; aCount < maxTimes; aCount++ {
		for bCount := 0; bCount < maxTimes; bCount++ {
			currX := m.buttonA.x*aCount + m.buttonB.x*bCount
			currY := m.buttonA.y*aCount + m.buttonB.y*bCount

			if currX == m.prizeLoc.x && currY == m.prizeLoc.y {
				return true, aCount, bCount
			}
		}
	}

	return false, 0, 0
}*/

func solveForMachine(m *machine, addBigNumber bool) (bool, int, int) {
	//I'm sure there's a better way to do this. Need to refresh my maths!

	//From hand solving the simultaneous equation:
	//ac+bd=p
	//ae+bf=q
	//Where...
	//a is the number of times A has to be pressed
	//b is the number of times B has to be pressed
	c := float64(m.buttonA.x)  //c is the x amount moved from pressing A
	d := float64(m.buttonB.x)  //d is the x amount moved from pressing B
	e := float64(m.buttonA.y)  //e is the y amount moved from pressing A
	f := float64(m.buttonB.y)  //f is the y amount moved from pressing B
	p := float64(m.prizeLoc.x) //p is the target x location
	q := float64(m.prizeLoc.y) //q is the target y location

	if addBigNumber {
		p += 10000000000000
		q += 10000000000000
	}

	denomB := (f*c - d*e)
	if denomB == 0 {
		panic("What?!")
	}
	b := (q*c - p*e) / denomB

	denomA := c
	if denomA == 0 {
		panic("What?!")
	}
	a := (p - b*d) / denomA

	aMod := math.Mod(a, 1)
	bMod := math.Mod(b, 1)

	return aMod == 0 && bMod == 0, int(a), int(b)
}

func runPuzzle(fileName string, addBigNumber bool) {
	machines := parseMachines(fileName)

	sumCost := 0

	for _, m := range machines {
		/*possible, numAPresses, numBPresses := testMachine(&m, 100)

		if possible {
			cost := numAPresses*3 + numBPresses
			sumCost += cost
		}*/

		possible, numAPresses, numBPresses := solveForMachine(&m, addBigNumber)

		if possible {
			cost := numAPresses*3 + numBPresses
			sumCost += cost
		}
	}

	fmt.Println("Sum Cost:", sumCost)
}

func Run(inputFileName string) {
	fmt.Println("Running Day 13 Puzzle 1...")
	runPuzzle(inputFileName, false)

	fmt.Println("Running Day 13 Puzzle 2...")
	runPuzzle(inputFileName, true)
}
