// Package day05
package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc-2025/helpers"
)

type rangeT struct {
	start, end int64
}

func Part1(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var ranges []rangeT
	var numbers []int64

	foundEmpty := false
	for _, line := range lines {
		if line == "" {
			foundEmpty = true
			continue
		}

		if !foundEmpty {
			a := strings.Split(line, "-")
			start, _ := strconv.ParseInt(a[0], 10, 64)
			end, _ := strconv.ParseInt(a[1], 10, 64)
			ranges = append(ranges, rangeT{start, end})
		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			numbers = append(numbers, num)
		}
	}

	slices.Sort(numbers)

	acc := int64(0)
	used := make(map[int64]bool)

	for _, num := range numbers {
		if used[num] {
			continue
		}
		for _, r := range ranges {
			if num >= r.start && num <= r.end {
				acc++
				used[num] = true
				break
			}
		}
	}

	return strconv.FormatInt(acc, 10), nil
}

func Part2(inputFile string) (string, error) {
	lines, err := helpers.ReadLines(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	var ranges []rangeT

	for _, line := range lines {
		if line == "" {
			break
		}
		a := strings.Split(line, "-")
		start, _ := strconv.ParseInt(a[0], 10, 64)
		end, _ := strconv.ParseInt(a[1], 10, 64)
		ranges = append(ranges, rangeT{start, end})
	}

	slices.SortFunc(ranges, func(a, b rangeT) int {
		if a.start < b.start {
			return -1
		}
		if a.start > b.start {
			return 1
		}
		return 0
	})

	var merged []rangeT
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[len(merged)-1]
		// Check if current range overlaps is next to the last merged range
		if r.start <= last.end+1 {
			// Extend if needed
			if r.end > last.end {
				last.end = r.end
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, r)
		}
	}

	// Calculate total count from merged ranges
	acc := int64(0)
	for _, r := range merged {
		acc += (r.end - r.start + 1)
	}

	return strconv.FormatInt(acc, 10), nil
}
