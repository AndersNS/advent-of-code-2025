// Package day04
package day04

import (
	"fmt"
	"strconv"

	"aoc-2025/helpers"
)

type Pos struct {
	X int
	Y int
}

type Cell struct {
	N int
	T rune
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	grid := make(map[Pos]Cell)
	acc := int64(0)

	directions := []Pos{
		{-1, -1}, {0, -1}, {1, -1}, // top
		{-1, 0}, {1, 0}, // same row
		{-1, 1}, {0, 1}, {1, 1}, // bottom
	}
	maxY := len(lines)
	maxX := len(lines[0])
	for y, str := range lines {
		for x, r := range str {
			// update own type
			val := grid[Pos{X: x, Y: y}]
			val.T = r
			grid[Pos{X: x, Y: y}] = val
			if r == '@' {
				for _, dir := range directions {
					newX, newY := x+dir.X, y+dir.Y
					if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY {
						neighbor := Pos{X: newX, Y: newY}
						val := grid[neighbor]
						val.N++
						grid[neighbor] = val
					}
				}
			}
		}
	}

	for _, v := range grid {
		if v.N < 4 && v.T == '@' {
			acc++
		}
	}

	return strconv.FormatInt(acc, 10), nil
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	blankGrid := make(map[Pos]Cell)
	removed := int64(0)

	directions := []Pos{
		{-1, -1}, {0, -1}, {1, -1}, // top
		{-1, 0}, {1, 0}, // same row
		{-1, 1}, {0, 1}, {1, 1}, // bottom
	}
	maxY := len(lines)
	maxX := len(lines[0])
	for y, str := range lines {
		for x, r := range str {
			val := blankGrid[Pos{X: x, Y: y}]
			val.T = r
			blankGrid[Pos{X: x, Y: y}] = val
		}
	}

	for {
		grid := CopyGrid(blankGrid)

		for y := 0; y < maxY; y += 1 {
			for x := 0; x < maxX; x += 1 {
				pos := Pos{X: x, Y: y}
				cell := grid[pos]
				if cell.T == '@' {
					for _, dir := range directions {
						newX, newY := x+dir.X, y+dir.Y
						if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY {
							neighbor := Pos{X: newX, Y: newY}
							val := grid[neighbor]
							val.N++
							grid[neighbor] = val
						}
					}
				}
			}
		}

		cellRemoved := false
		for y := 0; y < maxY; y += 1 {
			for x := 0; x < maxX; x += 1 {
				pos := Pos{X: x, Y: y}
				cell := grid[pos]
				if cell.N < 4 && cell.T == '@' {
					removed++
					cell.T = '.'
					blankGrid[pos] = cell
					cellRemoved = true
				}
			}
		}
		if !cellRemoved {
			return strconv.FormatInt(removed, 10), nil
		}
	}
}

func CopyGrid(original map[Pos]Cell) map[Pos]Cell {
	copy := make(map[Pos]Cell, len(original))

	for pos, value := range original {
		copy[pos] = value
	}

	return copy
}
