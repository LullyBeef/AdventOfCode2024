package day17

import (
	"aoc24/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type machine struct {
	registerA int
	registerB int
	registerC int
}

type opCode int

const (
	opAdv opCode = iota
	opBxl
	opBst
	opJnz
	opBxc
	opOut
	opBdv
	opCdv
	opMax
)

type program []int

func parseInput(fileName string) (machine, program) {
	var mach machine
	var prog program
	utils.ReadFileWithParse(fileName, func(line string) {
		if len(line) > 0 {
			switch {
			case strings.HasPrefix(line, "Register A: "):
				val, err := strconv.Atoi(strings.TrimPrefix(line, "Register A: "))
				if err != nil {
					panic(err)
				}
				mach.registerA = val
			case strings.HasPrefix(line, "Register B: "):
				val, err := strconv.Atoi(strings.TrimPrefix(line, "Register B: "))
				if err != nil {
					panic(err)
				}
				mach.registerB = val
			case strings.HasPrefix(line, "Register C: "):
				val, err := strconv.Atoi(strings.TrimPrefix(line, "Register C: "))
				if err != nil {
					panic(err)
				}
				mach.registerC = val
			case strings.HasPrefix(line, "Program: "):
				progStr := strings.Split(strings.TrimPrefix(line, "Program: "), ",")

				for i := 0; i < len(progStr); i += 2 {
					op, err := strconv.Atoi(progStr[i])
					if err != nil {
						panic(err)
					}

					operand, err := strconv.Atoi(progStr[i+1])
					if err != nil {
						panic(err)
					}

					prog = append(prog, op)
					prog = append(prog, operand)
				}
			default:
				panic("Unhandled!")
			}
		}
	})

	return mach, prog
}

func comboOperand(mach machine, operand int) int {
	switch operand {
	case 4:
		return mach.registerA
	case 5:
		return mach.registerB
	case 6:
		return mach.registerC
	case 7:
		panic("Apparently unused...")
	}

	return operand
}

func runProgram(mach machine, prog program, outputOrVerify bool) bool {
	instPtr := 0

	nextOutputIdx := 0

	progLen := len(prog)

	valid := true

	for instPtr = 0; instPtr < progLen; {
		op := opCode(prog[instPtr])
		operand := prog[instPtr+1]

		switch op {
		case opAdv:
			mach.registerA >>= operand
		case opBxl:
			mach.registerB ^= operand
		case opBst:
			mach.registerB = (comboOperand(mach, operand) % 8)
		case opJnz:
			if mach.registerA != 0 {
				instPtr = operand
				continue
			}
		case opBxc:
			mach.registerB ^= mach.registerC
		case opOut:
			if outputOrVerify {
				fmt.Print(comboOperand(mach, operand)%8, ",")
			} else {
				//fmt.Print(comboOperand(mach, operand)%8, ",")
				val := comboOperand(mach, operand) % 8
				if nextOutputIdx >= progLen {
					panic("Shouldn't happen!")
				} else if val != prog[nextOutputIdx] {
					//return false
					valid = false
					//break
				}
				nextOutputIdx++
			}

		case opBdv:
			mach.registerB = mach.registerA >> comboOperand(mach, operand)
		case opCdv:
			mach.registerC = mach.registerA >> comboOperand(mach, operand)
		default:
			panic("Unhandled!")
		}

		instPtr += 2
	}

	if valid {
		valid = nextOutputIdx == progLen
	}

	return valid
}

func runPuzzle1(fileName string) {
	mach, prog := parseInput(fileName)

	runProgram(mach, prog, true)
}

/*fmt.Printf("rega = 0x%x\n", rega)

fmt.Print("regb := rega % 8\n")
regb := rega % 8
fmt.Printf("\tregb is 0x%x\n", regb)

fmt.Print("regb ^= 0b101\n")
regb ^= 0b101
fmt.Printf("\tregb is 0x%x\n", regb)

fmt.Print("regc := rega >> regb\n")
regc := rega >> regb
fmt.Printf("\tregc is 0x%x\n", regc)

fmt.Print("regb ^= 0b110\n")
regb ^= 0b110
fmt.Printf("\tregb is 0x%x\n", regb)

fmt.Print("regb ^= regc\n")
regb ^= regc
fmt.Printf("\tregb is 0x%x\n", regb)

fmt.Print("out := regb % 8\n")
out := regb % 8
fmt.Printf("\tout is 0x%x\n", out)
//fmt.Printf("%d,", regb%8)
if nextOutputIdx >= progLen {
	panic("Shouldn't happen!")
} else if out != prog[nextOutputIdx] {
	valid = false
	//break
}
nextOutputIdx++

fmt.Print("rega >>= 3\n")
rega >>= 3
fmt.Printf("\trega is 0x%x\n", rega)
if rega == 0 {
	break
}*/

func runPuzzle2(fileName string) {
	//Example
	//117440 in binary, split into 8 octals: 011 100 101 011 000 000
	/*rega := 0b11100101011000000
	nextOutputIdx := 0
	for {
		//For each three bits of rega, spit out the number.
		rega >>= 3
		out := rega % 8
		fmt.Printf("%x,", out)
		nextOutputIdx++

		if rega == 0 {
			break
		}
	}*/

	_, prog := parseInput(fileName)

	progLen := len(prog)

	//startFrom := 0b011		//For just the last digit...
	//startFrom := 0b011000		//For the last two digits...
	//startFrom := 0b011000101	//For the last three digits... 101 gets xor'd with 110 to produce the desired 011 (3)
	//startFrom := 0b011000101001	//For the last four digits...
	//startFrom := 0b011000101001011	//For the last five digits...
	//startFrom := 0b011000101001011010	//For the last six digits...

	//startFrom := 0b011000101
	//endsWith := 0b11000101
	knownLSB := 0b11000101

	//startFrom := int(math.Pow(8, float64(progLen-1)))
	//startFrom := 64
	//startFrom := 0
	//startsWith := 1 << 9
	knownMSB := int(math.Pow(8, float64(progLen-1)))

	startFrom := knownMSB >> 8

	for i := startFrom; ; i++ {
		//fmt.Printf("%O:", i)

		rega := ((i << 8) | knownLSB)
		//fmt.Printf("%b\n", rega)

		nextOutputIdx := 0
		valid := true
		for {
			regb := rega % 8
			regb ^= 0b101
			regc := rega >> regb
			regb ^= 0b110
			regb ^= regc
			out := regb % 8
			//fmt.Printf("%d,", out)

			if nextOutputIdx >= progLen {
				panic("Shouldn't happen!")
			} else if out != prog[nextOutputIdx] {
				valid = false
				break
			}
			nextOutputIdx++

			rega >>= 3
			if rega == 0 {
				break
			}
		}

		//fmt.Print("\n")

		if nextOutputIdx == progLen {
			if valid {
				fmt.Println("Found at", i)
				break
			}
		} /*else {
			panic("Shouldn't happen!")
		}*/
	}

	fmt.Println("BLAH!")

	/*for i := 0; ; i++ {
		if runProgram(machine{i, 0, 0}, prog, false) {
			fmt.Println("Found at", i)
			break
		}
	}*/
}

func Run(inputFileName string) {
	//fmt.Println("Running Day 17 Puzzle 1...")
	//runPuzzle1(inputFileName)

	fmt.Println("Running Day 17 Puzzle 2...")
	runPuzzle2(inputFileName)
}
