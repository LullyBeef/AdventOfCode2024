package day8

import (
	"aoc24/utils"
	"fmt"
	"slices"
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

type antennaeMap map[rune][]vec2d

func findAntinodes(antennae []vec2d, gridWidth int, gridHeight int, useHarmonics bool) []vec2d {
	var antinodes []vec2d
	for a := range len(antennae) {
		for b := range len(antennae) {
			if a != b {
				if useHarmonics {
					antinodes = append(antinodes, antennae[a])
					antinodes = append(antinodes, antennae[b])
				}

				fromAntennae := antennae[b]

				diff := fromAntennae.sub(&antennae[a])

				for {
					antinode := fromAntennae.add(&diff)

					if antinode.x >= 0 &&
						antinode.y >= 0 &&
						antinode.x < gridWidth &&
						antinode.y < gridHeight {
						antinodes = append(antinodes, antinode)

						if useHarmonics {
							fromAntennae = antinode
						} else {
							break
						}
					} else {
						break
					}
				}
			}
		}
	}

	return antinodes
}

func runPuzzle(fileName string, useHarmonics bool) {
	gridWidth := 0
	gridHeight := 0
	antennae := make(antennaeMap)

	utils.ReadFileWithParse(fileName, func(line string) {
		for x, c := range line {
			if c != '.' {
				antennae[c] = append(antennae[c], vec2d{x, gridHeight})
			}
		}

		gridWidth = len(line)
		gridHeight++
	})

	var uniqueAntinodes []vec2d
	for _, m := range antennae {
		antinodes := findAntinodes(m, gridWidth, gridHeight, useHarmonics)

		for _, a := range antinodes {
			if !slices.Contains(uniqueAntinodes, a) {
				uniqueAntinodes = append(uniqueAntinodes, a)
			}
		}
	}

	fmt.Println("Unique Antinodes: ", len(uniqueAntinodes))
}

func Run(inputFileName string) {
	fmt.Println("Running Day 8 Puzzle 1...")
	runPuzzle(inputFileName, false)

	fmt.Println("Running Day 8 Puzzle 2...")
	runPuzzle(inputFileName, true)
}
