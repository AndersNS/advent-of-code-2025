// Package day07
package day07

import (
	"fmt"
	"strconv"

	"aoc-2025/helpers"
)

type Point struct {
	x int
	y int
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	grid := make(map[int][]string)
	maxY := len(lines)
	maxX := len(lines[0])

	visited := make(map[Point]bool)
	queue := helpers.Queue[Point]{}
	for y := 0; y < maxY; y++ {
		line := lines[y]
		for x := 0; x < maxX; x++ {
			val := string(line[x])
			if val == "S" {
				start := Point{x: x, y: y + 1}
				visited[start] = true
				queue.Enqueue(start)
				fmt.Printf("Found S %d %d", x, y)
			}
			grid[y] = append(grid[y], val)
		}
	}

	directions := []Point{
		{-1, 1}, {1, 1},
	}
	acc := int64(0)
	for !queue.IsEmpty() {
		curr, _ := queue.Dequeue()
		cell := grid[curr.y][curr.x]
		if cell == "^" {
			// Two new beams
			acc++
			for _, dir := range directions {
				newX, newY := curr.x+dir.x, curr.y
				if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY {
					next := Point{x: newX, y: newY}

					if !visited[next] {
						visited[next] = true
						queue.Enqueue(next)
					}
				}
			}
		} else {
			// We are going straight
			newX, newY := curr.x, curr.y+1
			next := Point{x: newX, y: newY}
			grid[curr.y][curr.x] = "|"
			if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY {
				if !visited[next] {
					visited[next] = true
					queue.Enqueue(next)
				}
			}
		}
	}

	return strconv.FormatInt(acc, 10), nil
}

func printGrid(grid map[int][]string, maxY int) {
	for y := 0; y <= maxY; y++ {
		g := grid[y]
		for _, v := range g {
			fmt.Printf("%s", v)
		}
		fmt.Println("")
	}
}

func countPaths(grid map[int][]string, p Point, maxX, maxY int, seen map[Point]int64) int64 {
	// Reached bottom
	if p.y >= maxY {
		return 1
	}

	// Path is invalid
	if p.x < 0 || p.x >= maxX {
		return 0
	}

	// Check if we have been here before
	if val, exists := seen[p]; exists {
		return val
	}

	// Get current cell
	cell := grid[p.y][p.x]
	var paths int64

	if cell == "^" {
		leftPaths := countPaths(grid, Point{x: p.x - 1, y: p.y + 1}, maxX, maxY, seen)
		rightPaths := countPaths(grid, Point{p.x + 1, p.y + 1}, maxX, maxY, seen)
		paths = leftPaths + rightPaths
	} else {
		// Go down
		paths = countPaths(grid, Point{x: p.x, y: p.y + 1}, maxX, maxY, seen)
	}

	// Store in memo
	seen[p] = paths
	return paths
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	grid := make(map[int][]string)
	maxY := len(lines)
	maxX := len(lines[0])

	var startX, startY int
	for y := 0; y < maxY; y++ {
		line := lines[y]
		for x := 0; x < maxX; x++ {
			val := string(line[x])
			if val == "S" {
				startX = x
				startY = y
			}
			grid[y] = append(grid[y], val)
		}
	}

	seen := make(map[Point]int64)
	totalPaths := countPaths(grid, Point{x: startX, y: startY}, maxX, maxY, seen)

	return strconv.FormatInt(totalPaths, 10), nil
}
