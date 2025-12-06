// Package day06
package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

type column struct {
	size  int
	value string
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var operations []string

	lastLine := lines[len(lines)-1]

	operations = append(operations, strings.Fields(lastLine)...)

	parts := strings.Fields(lines[0])

	var numbers []int64
	for _, part := range parts {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}

	// y goes down
	for y := 1; y < (len(lines) - 1); y++ {
		lower := strings.Fields(lines[y])
		// x goes to the right
		for x, n := range lower {
			num, _ := strconv.ParseInt(n, 10, 64)
			res := doMath(operations[x], numbers[x], num)

			numbers[x] = res
		}
	}

	acc := int64(0)

	for _, num := range numbers {
		acc += num
	}

	return strconv.FormatInt(acc, 10), nil
}

func doMath(op string, a int64, b int64) int64 {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	}

	fmt.Printf("unknown operation: %s", op)
	return 0
}

func parseColumns(line string) []column {
	re := regexp.MustCompile(`[*+]`)
	matches := re.FindAllStringIndex(line, -1)

	var columns []column
	for i, match := range matches {
		col := column{
			value: line[match[0]:match[1]],
		}
		if i < len(matches)-1 {
			col.size = matches[i+1][0] - match[0] - 1
		} else {
			col.size = len(line) - match[0]
		}
		columns = append(columns, col)
	}

	return columns
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	// Parse last line to get operations with their column sizes
	lastLine := lines[len(lines)-1]
	columns := parseColumns(lastLine)

	// Lets turn it around
	numberCols := make(map[int][]string)
	for y := 0; y < len(lines)-1; y++ {
		line := lines[y]
		p := 0
		for x := 0; x < len(columns); x++ {
			c := columns[x]
			column := line[p : p+c.size]
			p += c.size + 1
			for n := c.size - 1; n >= 0; n-- {
				val := column[n]
				// Create slice with c.size empty strings if it doesn't exist
				if _, exists := numberCols[x]; !exists {
					numberCols[x] = make([]string, c.size)
				}

				// Append to the string at index n
				numberCols[x][n] += string(val)
			}
		}
	}

	acc := int64(0)
	for y := 0; y < len(columns); y++ {
		number, _ := strconv.ParseInt(strings.TrimSpace(numberCols[y][0]), 10, 64)
		for x := 1; x < len(numberCols[y]); x++ {
			num, _ := strconv.ParseInt(strings.TrimSpace(numberCols[y][x]), 10, 64)
			number = doMath(columns[y].value, number, num)
		}

		acc += number
	}

	return strconv.FormatInt(acc, 10), nil
}
