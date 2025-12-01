// Package day01
package day01

import (
	"fmt"
	"strconv"

	"aoc-2025/helpers"
)

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	acc := 0
	pointer := 50
	for _, line := range lines {
		direction := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		switch direction {
		case 'L':
			pointer = wrap(pointer-number, 100)
		case 'R':
			pointer = wrap(pointer+number, 100)
		}

		if pointer == 0 {
			acc += 1
		}
	}

	return strconv.Itoa(acc), nil
}

func wrap(value, max int) int {
	return ((value % max) + max) % max
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	acc := 0
	pointer := 50
	for _, line := range lines {
		direction := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		nonWrapped := 0
		prevPointer := pointer
		crossings := 0

		switch direction {
		case 'L':
			nonWrapped = pointer - number
			// Negative number means we crossed 0
			if nonWrapped < 0 {
				// If we start at 0 we count full 100s
				if pointer == 0 {
					crossings = number / 100
				} else {
					// Otherwise we need to account for the first partial crossing
					crossings = (number-pointer)/100 + 1
				}
			}
		case 'R':
			nonWrapped = pointer + number
			if nonWrapped >= 100 {
				crossings = (nonWrapped) / 100
			}
		}

		pointer = wrap(nonWrapped, 100)

		if pointer == 0 && prevPointer != 0 && crossings == 0 {
			acc += 1
		} else {
			acc += crossings
		}
	}

	return strconv.Itoa(acc), nil
}
