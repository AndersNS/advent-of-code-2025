// Package day03
package day03

import (
	"fmt"
	"strconv"
	"sync"

	"aoc-2025/helpers"
)

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	acc := int64(0)
	for _, str := range lines {
		first := 0
		last := 0
	digit:
		for i, r := range str {
			digit := int(r - '0')
			if i < len(str)-1 {
				if digit > first {
					first = digit
					last = 0
					continue digit // Only use number once
				}
			}
			if i > 0 {
				if digit > last {
					last = digit
				}
			}
		}
		acc += int64((first * 10) + last)
	}
	return strconv.FormatInt(acc, 10), nil
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	results := make([]int64, len(lines))
	numWorkers := 6
	jobs := make(chan int, len(lines))
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobs {
				line := lines[index]
				batterySize := 12
				toRemove := len(line) - batterySize
				stack := make([]byte, 0, batterySize)

				for i := 0; i < len(line); i++ {
					// While we can still remove digits and current is bigger than stack top
					for len(stack) > 0 && toRemove > 0 && line[i] > stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
						toRemove--
					}
					stack = append(stack, line[i])
				}

				// Take first k digits and convert to number
				result := int64(0)
				for j := 0; j < batterySize; j++ {
					result = result*10 + int64(stack[j]-'0')
				}
				results[index] = result
			}
		}()
	}

	// Send jobs to workers
	for i := 0; i < len(lines); i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	acc := int64(0)
	for _, result := range results {
		acc += result
	}

	return strconv.FormatInt(acc, 10), nil
}
