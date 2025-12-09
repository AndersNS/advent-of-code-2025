// Package day09
package day09

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

type Point struct {
	X int64
	Y int64
}

type PointPair struct {
	p1   Point
	p2   Point
	area int64
}

func (p1 Point) Area(p2 Point) int64 {
	xDiff := int64(0)
	if p1.X > p2.X {
		xDiff = p1.X - p2.X
	} else {
		xDiff = p2.X - p1.X
	}

	yDiff := int64(0)
	if p1.Y > p2.Y {
		yDiff = p1.Y - p2.Y
	} else {
		yDiff = p2.Y - p1.Y
	}

	return (xDiff + 1) * (yDiff + 1)
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var points []Point

	for _, line := range lines {
		a := strings.Split(line, ",")
		x, _ := strconv.ParseInt(a[0], 10, 64)
		y, _ := strconv.ParseInt(a[1], 10, 64)
		point := Point{X: x, Y: y}

		points = append(points, point)
	}

	biggestArea := int64(0)
	for i, p1 := range points {
		for j, p2 := range points {
			if j == i {
				continue
			}

			area := p1.Area(p2)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}

	return strconv.FormatInt(biggestArea, 10), nil
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var redpoints []Point
	completeGrid := make(map[Point]string)
	redpointsY := make(map[int64][]Point)
	redpointsX := make(map[int64][]Point)

	maxX := int64(0)
	maxY := int64(0)

	for _, line := range lines {
		a := strings.Split(line, ",")
		x, _ := strconv.ParseInt(a[0], 10, 64)
		y, _ := strconv.ParseInt(a[1], 10, 64)
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}

		point := Point{X: x, Y: y}

		redpointsY[y] = append(redpointsY[y], point)
		redpointsX[x] = append(redpointsX[x], point)
		completeGrid[point] = "r"
		redpoints = append(redpoints, point)
	}

	max := int64(math.Max(float64(maxX), float64(maxY)))

	greenpointsY := make(map[int64][]Point)
	greenpointsX := make(map[int64][]Point)
	for i := int64(0); i <= max; i++ {
		// First add greenpoints on x axis
		if xpoints, exists := redpointsX[i]; exists {
			lowY := findLowestY(xpoints).Y
			highY := findHighestY(xpoints).Y

			for y := lowY + 1; y < highY; y++ {
				greenpoint := Point{X: i, Y: y}
				greenpointsX[i] = append(greenpointsX[i], greenpoint)
				greenpointsY[y] = append(greenpointsY[y], greenpoint)
				completeGrid[greenpoint] = "g"
			}
		}
		if ypoints, exists := redpointsY[i]; exists {
			lowX := findLowestX(ypoints).X
			highX := findHighestX(ypoints).X

			for x := lowX + 1; x < highX; x++ {
				greenpoint := Point{X: x, Y: i}
				greenpointsX[x] = append(greenpointsX[x], greenpoint)
				greenpointsY[i] = append(greenpointsY[i], greenpoint)
				completeGrid[greenpoint] = "g"
			}
		}
		// Then add greenpoints on y axis
	}

	// Then draw green ones inside green
	for i := int64(0); i <= max; i++ {
		if xpoints, exists := greenpointsX[i]; exists {
			lowY := findLowestY(xpoints).Y
			highY := findHighestY(xpoints).Y

			for y := lowY + 1; y < highY; y++ {
				greenpoint := Point{X: i, Y: y}
				if _, exists := completeGrid[greenpoint]; !exists {
					greenpointsX[i] = append(greenpointsX[i], greenpoint)
					greenpointsY[y] = append(greenpointsY[y], greenpoint)
					completeGrid[greenpoint] = "g"
				}
			}
		}
		if ypoints, exists := greenpointsY[i]; exists {
			lowX := findLowestX(ypoints).X
			highX := findHighestX(ypoints).X

			for x := lowX + 1; x < highX; x++ {
				greenpoint := Point{X: x, Y: i}
				if _, exists := completeGrid[greenpoint]; !exists {
					greenpointsX[x] = append(greenpointsX[x], greenpoint)
					greenpointsY[i] = append(greenpointsY[i], greenpoint)
					completeGrid[greenpoint] = "g"
				}
			}
		}
	}

	// printGrid(completeGrid, maxY, maxX)

	// Generate all point pairs and sort by area (descending)
	var pairs []PointPair
	for i, p1 := range redpoints {
		for j, p2 := range redpoints {
			if j <= i {
				continue
			}
			area := p1.Area(p2)
			pairs = append(pairs, PointPair{p1: p1, p2: p2, area: area})
		}
	}

	// Sort pairs by area in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area > pairs[j].area
	})

	// Check largest rectangles first, with early termination
	biggestArea := int64(0)
	for _, pair := range pairs {
		// Early termination - if remaining pairs can't beat current best, stop
		if pair.area <= biggestArea {
			break
		}
		if isRectangleFilled(pair.p1, pair.p2, completeGrid) {
			biggestArea = pair.area
		}
	}

	return strconv.FormatInt(biggestArea, 10), nil
}

func isRectangleFilled(p1, p2 Point, grid map[Point]string) bool {
	minX := p1.X
	maxX := p2.X
	if p2.X < p1.X {
		minX = p2.X
		maxX = p1.X
	}

	minY := p1.Y
	maxY := p2.Y
	if p2.Y < p1.Y {
		minY = p2.Y
		maxY = p1.Y
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, exists := grid[Point{X: x, Y: y}]; !exists {
				return false
			}
		}
	}
	return true
}

func printGrid(grid map[Point]string, maxY, maxX int64) {
	if len(grid) == 0 {
		return
	}
	for y := int64(0); y <= maxY; y++ {
		fmt.Printf("%d ", y)
		for x := int64(0); x <= maxX; x++ {
			if val, ok := grid[Point{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findHighestY(points []Point) Point {
	if len(points) == 0 {
		panic("empty points list")
	}

	highPoint := points[0]
	for _, p := range points[1:] {
		if p.Y > highPoint.Y {
			highPoint = p
		}
	}
	return highPoint
}

func findLowestY(points []Point) Point {
	if len(points) == 0 {
		panic("empty points list")
	}

	minPoint := points[0]
	for _, p := range points[1:] {
		if p.Y < minPoint.Y {
			minPoint = p
		}
	}
	return minPoint
}

func findHighestX(points []Point) Point {
	if len(points) == 0 {
		panic("empty points list")
	}

	highPoint := points[0]
	for _, p := range points[1:] {
		if p.X > highPoint.X {
			highPoint = p
		}
	}
	return highPoint
}

func findLowestX(points []Point) Point {
	if len(points) == 0 {
		panic("empty points list")
	}

	minPoint := points[0]
	for _, p := range points[1:] {
		if p.X < minPoint.X {
			minPoint = p
		}
	}
	return minPoint
}
